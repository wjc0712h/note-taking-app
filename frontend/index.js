loadNotes()

const saveBtn = document.getElementById("save-button")

async function login() {
    const username = document.getElementById("username");
    const res = await window.api.login(username.value);

    res.success ? (console.log("Logged in as", username.value), loadNotes(),username.value = "") : createProfile(username.value);
    console.log(res.message);
}

async function createProfile(username) {
    const res = await window.api.createProfile(username);
    res.success ? (console.log("Profile created"), login(username)) : console.log("Profile not created");
    console.log(res.message);
}
async function loadNotes() {
    const notes = await window.api.fetchNotes();
    const notesList = document.getElementById("notes-list");
    notesList.innerHTML = "";

    const profile = await window.api.fetchProfile();
    const user = document.getElementById("profile");
    if (profile == null) return; 

    user.innerHTML = profile.username;
    if (notes != null) {
        notes.forEach(note => addNoteToList(note));
    }
}
function addNoteToList(note) {
    const notesList = document.getElementById("notes-list");

    const li = document.createElement("li");

    li.textContent = note.content?.length > 15 ? note.content.slice(0, 15) + "..." : note.content || "";
    li.setAttribute("data-id", note.id);

    li.addEventListener("click", () => {
        document.getElementById("note-content").value = note.content;
        document.getElementById("note-title").innerText = note.createdAt;
        saveBtn.textContent = "UPDATE";
        saveBtn.className = note.id;

        const existingDeleteBtn = document.getElementById("delete-button");
        if (existingDeleteBtn) {
            existingDeleteBtn.remove();
        }

        const deleteBtn = document.createElement("button");
        deleteBtn.textContent = "DELETE";
        deleteBtn.id = "delete-button";
        document.getElementById("buttons").appendChild(deleteBtn);

        deleteBtn.addEventListener("click", deleteNote);
    });

    notesList.appendChild(li);
}
async function saveNote() {
    const content = document.getElementById("note-content");
    
    if (!content.value.trim()) return; 

    if (saveBtn.className) {
        await window.api.updateNote(saveBtn.className, content.value);

        const selectedNote = document.querySelector(`li[data-id="${saveBtn.className}"]`);
        if (selectedNote) {
            selectedNote.textContent = content.value?.length > 15 ? content.value.slice(0, 15) + "..." : content.value || "";
        }
    } else {
        const newNote = await window.api.createNote(content.value);
        if (newNote) {
            addNoteToList(newNote);
            content.value = "";
        }
    }
    loadNotes()
}
async function deleteNote() {
    const noteId = saveBtn.className;
    if (!noteId) return;

    try {
        const response = await window.api.deleteNote(noteId);
        saveBtn.className = ""
        saveBtn.textContent = "SAVE";
        const existingDeleteBtn = document.getElementById("delete-button");
        if (existingDeleteBtn) {
            existingDeleteBtn.remove();
        }
        document.getElementById("note-content").value = "";
        document.getElementById("note-title").innerHTML = "New Note"
        if (!response || response.error) {
            console.error("Failed to delete note:", response?.error || "Unknown error");
            return;
        }

        loadNotes();
    } catch (error) {
        console.error("Error deleting note:", error);
    }
}