const Estatus = {
    Completada: "Completada",
    Pendiente: "Pendiente"
}

async function LoadParams() {
    const response = await fetch('http://localhost:8080/tasks')
    if (!response.ok) throw new Error(`Error HTTP: ${response.status}`)
    const tasks = await response.json();
    renderTask(tasks)
}

function renderTask(tasks){
    const taskList = document.getElementById('task-list');
    taskList.innerHTML = '';
    tasks.forEach(task => {
        const taskItem = document.createElement('li');
        taskItem.className = 'task-item';
        taskItem.innerHTML = `
            <div class="task-content">
                <h3>${task.title}</h3>
                <p>${task.description}</p>
                <small>Estado: ${task.is_done ? Estatus.Completada : Estatus.Pendiente}</small>
            </div>
            <div class="task-actions">
                <button class="edit" onclick="modalEdit(${task.id}, '${task.title}', '${task.description}', ${task.is_done})">Editar</button>
                <button class="delete" onclick="DeleteTask(${task.id})">Eliminar</button>
            </div>
        `;
        taskList.appendChild(taskItem);
    });
}

async function CreatTask(event) {
    event.preventDefault(); // Evitar que el formulario se envíe y recargue la página
    const title = document.getElementById('task-title').value;
    const description = document.getElementById('task-description').value;
    const task = {
        title: title,
        description: description,
        is_done: false
    }

    const response = await fetch('http://localhost:8080/tasks', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(task)
    });

    if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);
    const newTask = await response.json();
    LoadParams();
}

async function DeleteTask(id) {
    const response = await fetch(`http://localhost:8080/tasks/${id}`, {
        method: 'DELETE'
    });

    if (!response.ok) {
        throw new Error(`Error HTTP: ${response.status}`);
    }

    LoadParams(); // Recargar la lista de tareas
}

async function EditTask(id) {
    const title = document.getElementById('edit-task-title').value;
    const description = document.getElementById('edit-task-description').value;
    const is_done = document.getElementById('edit-task-is_done').value === 'true';
    const task = {
        title: title,
        description: description,
        is_done: is_done
    };

    const response = await fetch(`http://localhost:8080/tasks/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(task)
    });

    if (!response.ok) {
        throw new Error(`Error HTTP: ${response.status}`);
    }

    const updatedTask = await response.json();
    LoadParams(); 
    closeModal(); 
}

function modalEdit(id, title, description, is_done) {
    document.getElementById('edit-task-title').value = title;
    document.getElementById('edit-task-description').value = description;
    document.getElementById('edit-task-is_done').value = is_done;
    document.getElementById('save-task-button').onclick = function() {
        EditTask(id);
    };
    document.getElementById('edit-modal').style.display = 'block';
}

function closeModal() {
    document.getElementById('edit-modal').style.display = 'none';
}

document.addEventListener('DOMContentLoaded', () => {
    LoadParams();

    document.getElementById('create-task-button').addEventListener('click', (event) => {
        CreatTask(event).catch(error => console.error('Error:', error));
    });
});