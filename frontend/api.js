const { ipcMain, session } = require("electron");
const axios = require("axios");

const apiRequest = async (method, endpoint, data = null) => {
  try {
    const cookies = (await session.defaultSession.cookies.get({ url: "http://localhost:8080" }))
      .map(({ name, value }) => `${name}=${value}`).join("; ");
    const headers = cookies ? { Cookie: cookies } : {};
    const config = { method, url: `http://localhost:8080/api/${endpoint}`, headers, withCredentials: true, data };
    const response = await axios(config);
    return response.data;
  } catch (error) {
    console.error(`Error in ${endpoint}:`, error.response?.data || error.message);
    return null;
  }
};

ipcMain.handle("login", async (_, username) => {
  try {
    const response = await axios.post("http://localhost:8080/api/auth/login", { username }, { withCredentials: true });
    const cookies = response.headers["set-cookie"] || [];
    for (const cookie of cookies) {
      const [key, value] = cookie.split(";")[0].split("=");
      await session.defaultSession.cookies.set({ url: "http://localhost:8080", name: key.trim(), value: value.trim() });
    }
    return { success: true, message: response.data };
  } catch (error) {
    console.error("Login failed:", error.response?.data || error.message);
    return { success: false, message: "Login failed" };
  }
});

ipcMain.handle("create-profile", (_, username) => apiRequest("post", "profile/new", { username }));
ipcMain.handle("fetch-profile", () => apiRequest("get", "profile/me"));

ipcMain.handle("fetch-note", () => apiRequest("get", "note/all"));
ipcMain.handle("create-note", (_, content) => apiRequest("put", "note/create", { content }));
ipcMain.handle("update-note", (_, id, content) => apiRequest("patch", `note/update/${id}`, { content }));
ipcMain.handle("delete-note", (_, id) => apiRequest("delete", `note/delete/${id}`));
