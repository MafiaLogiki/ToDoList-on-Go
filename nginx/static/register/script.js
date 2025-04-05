document.getElementById('registrationForm').addEventListener('submit', function(e) {
    e.preventDefault();
    
    // Очищаем сообщения
    document.getElementById('errorMessage').textContent = '';
    document.getElementById('successMessage').textContent = '';
    
    // Получаем значения полей
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    
    // Проверка совпадения паролей
    if (password !== confirmPassword) {
        document.getElementById('errorMessage').textContent = 'Пароли не совпадают';
        return;
    }
    
    // Создаем объект с данными пользователя (без email)
    const userData = {
        username: username,
        password: password
    };
    
    // Отправляем запрос на сервер
    fetch('/api/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData)
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(err => { throw err; });
        }
        return response.json();
    })
    .then(data => {
        // Успешная регистрация
        document.getElementById('successMessage').textContent = 'Регистрация прошла успешно!';
        document.getElementById('registrationForm').reset();
    })
    .catch(error => {
        // Обработка ошибок
        let errorMessage = 'Произошла ошибка при регистрации';
        if (error.message) {
            errorMessage = error.message;
        }
        document.getElementById('errorMessage').textContent = errorMessage;
    });
});
