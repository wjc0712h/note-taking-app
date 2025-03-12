const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("api", {
  login: (username) => ipcRenderer.invoke("login", username),
  fetchNotes: () => ipcRenderer.invoke("fetch-notes"),
  createNote: (content) => ipcRenderer.invoke("create-note", content),
});

