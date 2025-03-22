const { app, BrowserWindow, session } = require("electron");
const path = require("path");
const { spawn } = require("child_process");
require('./api')

const fs = require("fs");

let mainWindow;
let serverProcess

//const logFile = path.join(app.getPath("home"), "app.log");
//const logStream = fs.createWriteStream(logFile, { flags: "a" });
app.whenReady().then(() => {

  const dbPath = path.join(app.getPath("home"), "db/database.db");
  serverProcess = spawn(path.join(__dirname, 'dist/server'), [dbPath]);

  // console.log(path.join(app.getPath("home"), "db/database.db"))
  // console.log(path.join(app.getPath("userData"), "app.log"))
  serverProcess.stdout.on("data", (data) => {
    const message = `[Go Server]: ${data}`;
    console.log(message);
    //logStream.write(message);
  });

  serverProcess.stderr.on("data", (data) => {
    const message = `[Go Server Error]: ${data}`;
    console.error(message);
    //logStream.write(message);
  });

  serverProcess.on("exit", (code, signal) => {
    const message = `Go Server exited with code ${code}, signal ${signal}\n`;
    console.log(message);
    //logStream.write(message);
  });
  mainWindow = new BrowserWindow({
    width: 700,
    height: 600,
    icon: path.join(__dirname, "icon.icns"),
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
      nodeIntegration: false,
      contextIsolation: true,
      session: session.defaultSession,
    },
  });
  mainWindow.setWindowButtonVisibility(true);
  mainWindow.loadFile("index.html");
});

app.on("window-all-closed", () => {
  if (serverProcess) {
    serverProcess.kill();
    console.log("Go server terminated.");
    //logStream.write("Go server terminated.\n");
  }
  app.quit();
});

app.on("before-quit", async () => {
  try {
    await session.defaultSession.clearStorageData({ storages: ["cookies"] });
    console.log("Cookies cleared before exit.");
  } catch (error) {
    console.error("Error clearing cookies:", error);
  } finally {
    //logStream.end();
  }
});