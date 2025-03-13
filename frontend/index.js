loadNotes()
async function login() {
    const username = document.getElementById("username").value;
    const response = await window.api.login(username);
    if (response.success) {
        console.log("successfully logged in as ", username)
        loadNotes();
    } else {
        console.log(response.message);
    }
}

async function loadNotes() {
    const notes = await window.api.fetchNotes();
    const notesList = document.getElementById("notes-list");
    notesList.innerHTML = "";

    const profile = await window.api.fetchProfile();
    const user = document.getElementById("profile");
    if(profile.username != undefined) {
        // const login = document.getElementById("login");
        user.innerHTML = profile.username
    }


    notes.forEach(note => {
        const li = document.createElement("li");
        li.textContent = note.content;
        notesList.appendChild(li);
    });
}

async function createNote() {
    const content = document.getElementById("note-content").value;
    await window.api.createNote(content);
    document.getElementById("note-content").value = "";
    loadNotes();
}