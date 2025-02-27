document.addEventListener("DOMContentLoaded", function () {
    // Инициализация переменных
    let menuButton = document.querySelector(".menu-toggle");
    let menu = document.querySelector(".sidebar");

    // Функция для открытия/закрытия меню
    function toggleMenu() {
        menu.classList.toggle("open");  // Открытие/закрытие меню

        let menuWidth = menu.offsetWidth;

        if (menu.classList.contains("open")) {
            menuButton.style.transform = `translateX(${menuWidth}px)`; // Сдвиг кнопки
        } else {
            menuButton.style.transform = "translateX(0)";
        }
    }

    // Открытие/закрытие по кнопке-бургеру
    if (menuButton) {
        menuButton.addEventListener("click", function (event) {
            event.stopPropagation();  // Предотвращаем всплытие, чтобы клик не сработал на документе
            toggleMenu();
        });
    }

    // Закрытие меню при клике вне его области
    document.addEventListener("click", function (event) {
        if (!menu.contains(event.target) && !menuButton.contains(event.target)) {
            menu.classList.remove("open");
            menuButton.style.transform = "translateX(0)"; // Возвращаем кнопку на место
        }
    });

    // Отключаем закрытие, если кликнули внутри меню
    if (menu) {
        menu.addEventListener("click", function (event) {
            event.stopPropagation();  // Предотвращаем всплытие
        });
    }
});

document.getElementById("registerForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Отменяем стандартную отправку формы

    let valid = true;

    // Очистим все предыдущие ошибки
    document.querySelectorAll('.error-message').forEach(function(errorDiv) {
        errorDiv.style.display = 'none';
    });
    document.querySelectorAll('input').forEach(function(input) {
        input.classList.remove('error');
    });

    // Валидация имени пользователя
    const username = document.getElementById("username").value;
    if (username.trim() === "") {
        document.getElementById("username-error").style.display = "block";
        document.getElementById("username-error").innerText = "Имя пользователя не может быть пустым";
        document.getElementById("username").classList.add("error");
        valid = false;
    }

    // Валидация email
    const email = document.getElementById("email").value;
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (!emailRegex.test(email)) {
        document.getElementById("email-error").style.display = "block";
        document.getElementById("email-error").innerText = "Некорректный email";
        document.getElementById("email").classList.add("error");
        valid = false;
    }

    // Валидация пароля
    const password = document.getElementById("password").value;
    if (password.length < 6) {
        document.getElementById("password-error").style.display = "block";
        document.getElementById("password-error").innerText = "Пароль должен содержать хотя бы 6 символов";
        document.getElementById("password").classList.add("error");
        valid = false;
    }

    // Валидация подтверждения пароля
    const confirmPassword = document.getElementById("confirm-password").value;
    if (confirmPassword !== password) {
        document.getElementById("confirm-password-error").style.display = "block";
        document.getElementById("confirm-password-error").innerText = "Пароли не совпадают";
        document.getElementById("confirm-password").classList.add("error");
        valid = false;
    }

    // Если все валидации прошли, отправляем форму
    if (valid) {
        document.getElementById("registerForm").submit();
    }
});

// Прячем ошибки при клике вне поля ввода
document.querySelectorAll('input').forEach(function(input) {
    input.addEventListener('focus', function() {
        document.getElementById(input.id + "-error").style.display = "none";
        input.classList.remove('error');
    });
});