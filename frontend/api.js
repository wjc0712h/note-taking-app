const { ipcMain, session } = require("electron");
const axios = require("axios");


ipcMain.handle("login", async (_, username) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/api/auth/login",
        { username },
        { withCredentials: true }
      );
  
      console.log("Login Response Headers:", response.headers);
  
      const setCookieHeader = response.headers["set-cookie"];
      if (setCookieHeader) {
        for (const cookie of setCookieHeader) {
          const parsedCookie = cookie.split(";")[0];
          const [key, value] = parsedCookie.split("=");
  
          await session.defaultSession.cookies.set({
            url: "http://localhost:8080",
            name: key.trim(),
            value: value.trim(),
            // expirationDate: Date.now() / 1000 + 3600, // 1 hour
          });
        }
      }
      console.log("Cookies After Login:", await session.defaultSession.cookies.get({}));
      return { success: true, message: response.data };
    } catch (error) {
      console.error("Login failed:", error.response?.data || error.message);
      return { success: false, message: "Login failed" };
    }
  });
  
  ipcMain.handle("fetch-notes", async () => {
    try {
      const cookies = await session.defaultSession.cookies.get({ url: "http://localhost:8080" });
      const cookieString = cookies.map(cookie => `${cookie.name}=${cookie.value}`).join("; ");
  
      const response = await axios.get("http://localhost:8080/api/note/all", {
        headers: { Cookie: cookieString },
        withCredentials: true,
      });
  
      console.log("Notes Response Headers:", response.headers);
      return response.data;
    } catch (error) {
      console.error("Error fetching notes:", error.response?.data || error.message);
      return [];
    }
  });
  
  
  ipcMain.handle("create-note", async (_, content) => {
    try {
      const cookies = await session.defaultSession.cookies.get({ url: "http://localhost:8080" });
      const cookieString = cookies.map(cookie => `${cookie.name}=${cookie.value}`).join("; ");
  
      const response = await axios.put(
        "http://localhost:8080/api/note/",
        { content },
        {  headers: { Cookie: cookieString },
        withCredentials: true }
      );
      return response.data;
    } catch (error) {
      console.error("Error creating note:", error);
      return null;
    }
  });