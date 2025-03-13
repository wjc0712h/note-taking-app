const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("api", {
  //auth
  login: (username) => ipcRenderer.invoke("login", username),

  //profile
  fetchProfile: () => ipcRenderer.invoke("fetch-profile"),

  //note
  fetchNotes: () => ipcRenderer.invoke("fetch-note"),
  createNote: (content) => ipcRenderer.invoke("create-note", content),
});

