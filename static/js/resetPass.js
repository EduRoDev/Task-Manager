document.addEventListener('DOMContentLoaded', () => {
    document.getElementById('resetPasswordForm').addEventListener('submit', (e) => {
        e.preventDefault();
        const new_password = document.getElementById('new_password').value;
        const email = document.getElementById('email').value;
        const token = document.getElementById('token').value;
        const data = {new_password, email, token};

        fetch('http://localhost:8080/user/ResetPass',{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
        .then(async response => {
            if (!response.ok) {
                const data = await response.json();
                alert(data.error);
                throw new Error(data.error);
            }
            return response.json();
        })
        .then(data => {
            console.log('Respuesta completa del servidor:', data);
            const successLink = document.getElementById('successLink');
            successLink.click();
            
        })
        .catch(error => {
            console.error('Error:', error);
        });
    })
})