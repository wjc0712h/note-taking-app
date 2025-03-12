const { app, BrowserWindow, session } = require("electron");
const path = require("path");
require("./api");

let mainWindow;

app.whenReady().then(() => {
  mainWindow = new BrowserWindow({
    width: 900,
    height: 900,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
      nodeIntegration: false,
      contextIsolation: true,
      session: session.defaultSession,
    },
  });

  mainWindow.loadFile("index.html");
});

app.on("before-quit", async () => {
  try {
    await session.defaultSession.clearStorageData({ storages: ["cookies"] });
    console.log("Cookies cleared before exit.");
  } catch (error) {
    console.error("Error clearing cookies:", error);
  }
});