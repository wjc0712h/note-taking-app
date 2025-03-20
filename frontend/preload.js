const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("api", {
  //auth
  login: (username) => ipcRenderer.invoke("login", username),

  //profile
  fetchProfile: () => ipcRenderer.invoke("fetch-profile"),
  createProfile: (username) => ipcRenderer.invoke("create-profile", username),

  //note
  fetchNotes: () => ipcRenderer.invoke("fetch-note"),
  createNote: (content) => ipcRenderer.invoke("create-note", content),
  updateNote: (id, content) => ipcRenderer.invoke("update-note",id,content),
  deleteNote: (_,id) => ipcRenderer.invoke("delete-note",_,id)
});

