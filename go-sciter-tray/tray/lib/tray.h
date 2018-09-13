#ifndef TRAY_H
#define TRAY_H

#include <windows.h>

#include <shellapi.h>

struct tray {
  char *icon;
};

static void tray_update(struct tray *tray);

#define WM_TRAY_CALLBACK_MESSAGE (WM_USER + 1)
#define WC_TRAY_CLASS_NAME "TRAY"
#define ID_TRAY_FIRST 1000

static WNDCLASSEX wc;
static NOTIFYICONDATA nid;
static HWND hwnd;

static LRESULT CALLBACK _tray_wnd_proc(HWND hwnd, UINT msg, WPARAM wparam, LPARAM lparam) {
  switch (msg) {
  case WM_CLOSE:
    DestroyWindow(hwnd);
    return 0;
  case WM_DESTROY:
    PostQuitMessage(0);
    return 0;
  case WM_TRAY_CALLBACK_MESSAGE:
    if (lparam == WM_LBUTTONUP || lparam == WM_RBUTTONUP) {
      POINT p;
      GetCursorPos(&p);

      if (lparam == WM_LBUTTONUP) {
        clickCallback(p.x, p.y, FALSE);
      }
      if (lparam == WM_RBUTTONUP) {
        clickCallback(p.x, p.y, TRUE);
      }

      return 0;
    }
    break;
  }
  return DefWindowProc(hwnd, msg, wparam, lparam);
}

static int tray_init(struct tray *tray) {
  memset(&wc, 0, sizeof(wc));
  wc.cbSize = sizeof(WNDCLASSEX);
  wc.lpfnWndProc = _tray_wnd_proc;
  wc.hInstance = GetModuleHandle(NULL);
  wc.lpszClassName = WC_TRAY_CLASS_NAME;
  if (!RegisterClassEx(&wc)) {
    return -1;
  }

  hwnd = CreateWindowEx(0, WC_TRAY_CLASS_NAME, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0);
  if (hwnd == NULL) {
    return -1;
  }
  UpdateWindow(hwnd);

  memset(&nid, 0, sizeof(nid));
  nid.cbSize = sizeof(NOTIFYICONDATA);
  nid.hWnd = hwnd;
  nid.uID = 0;
  nid.uFlags = NIF_ICON | NIF_MESSAGE;
  nid.uCallbackMessage = WM_TRAY_CALLBACK_MESSAGE;
  Shell_NotifyIcon(NIM_ADD, &nid);

  tray_update(tray);
  return 0;
}

static int tray_loop(int blocking) {
  MSG msg;
  if (blocking) {
    GetMessage(&msg, NULL, 0, 0);
  } else {
    PeekMessage(&msg, NULL, 0, 0, PM_REMOVE);
  }
  if (msg.message == WM_QUIT) {
    return -1;
  }
  TranslateMessage(&msg);
  DispatchMessage(&msg);
  return 0;
}

static void tray_update(struct tray *tray) {
  UINT id = ID_TRAY_FIRST;
  HICON icon;
  HINSTANCE hInst;

  static const int icon_size = 32;
  int offset = LookupIconIdFromDirectoryEx((PBYTE)tray->icon, TRUE, icon_size, icon_size, LR_DEFAULTCOLOR);
  if (offset != 0) {
    icon = CreateIconFromResourceEx((PBYTE)tray->icon + offset, 0, TRUE, 0x30000, icon_size, icon_size, LR_DEFAULTCOLOR);
  }

  if (nid.hIcon) {
    DestroyIcon(nid.hIcon);
  }
  nid.hIcon = icon;
  Shell_NotifyIcon(NIM_MODIFY, &nid);
}

static void tray_exit() {
  Shell_NotifyIcon(NIM_DELETE, &nid);
  if (nid.hIcon != 0) {
    DestroyIcon(nid.hIcon);
  }
  PostQuitMessage(0);
  UnregisterClass(WC_TRAY_CLASS_NAME, GetModuleHandle(NULL));
}

#endif /* TRAY_H */