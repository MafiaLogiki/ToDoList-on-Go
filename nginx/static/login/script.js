document.getElementById('registrationForm').addEventListener('submit', async function(event) {
    event.preventDefault();
    
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    const errorMessage = document.getElementById('errorMessage');
    const successMessage = document.getElementById('successMessage');
    
    // Очищаем предыдущие сообщения
    errorMessage.textContent = '';
    successMessage.textContent = '';
    
    // Валидация паролей
    if (password !== confirmPassword) {
        errorMessage.textContent = 'Пароли не совпадают';
        return;
    }
    
    try {
        // Хеширование пароля (используем встроенный API)
        const hashedPassword = await hashPassword(password);
        
        // Отправка данных на сервер
        const response = await fetch('http://auth-service:8081/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                password: hashedPassword
            })
        });
        
        if (response.ok) {
            const data = await response.json();
            successMessage.textContent = 'Регистрация успешна!';
            // Перенаправление на страницу входа через 2 секунды
            setTimeout(() => {
                window.location.href = '/login';
            }, 2000);
        } else {
            const errorData = await response.json();
            errorMessage.textContent = errorData.error || 'Ошибка регистрации';
        }
    } catch (error) {
        errorMessage.textContent = 'Ошибка соединения с сервером';
        console.error('Registration error:', error);
    }
});

// Обработчик кнопки входа
document.getElementById('loginBtn').addEventListener('click', function() {
    window.location.href = '/login';
});

// Функция для хеширования пароля
async function hashPassword(password) {
    const encoder = new TextEncoder();
    const data = encoder.encode(password);
    const hashBuffer = await crypto.subtle.digest('SHA-256', data);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(byte => byte.toString(16).padStart(2, '0')).join('');
}
