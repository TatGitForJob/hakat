// Получаем элементы с помощью метода querySelectorAll
const wrapper = document.querySelectorAll(".wrapper");
const iconMenu = document.querySelectorAll(".icon-menu");
const menuBody = document.querySelectorAll(".menu__body");

// Добавляем класс "loaded" к элементу "wrapper"
wrapper.forEach(function (item) {
  item.classList.add("loaded");
});

// Навешиваем обработчик событий на элемент "icon-menu"
iconMenu.forEach(function (item) {
  item.addEventListener("click", function (event) {
    // Переключаем класс "active" у элемента "icon-menu"
    this.classList.toggle("active");
    menuBody.forEach(function (menu) {
      menu.classList.toggle("active");
    });
    // Переключаем класс "lock" у элемента "body"
    document.body.classList.toggle("lock");
  });
});

// График
const ctx = document.getElementById("myChart");

new Chart(ctx, {
  type: "bar",
  data: {
    labels: [1, 2, 3, 4, 5, 6],
    datasets: [
      {
        label: "График динамики бронирования",
        data: [12, 19, 3, 5, 2, 3],
        borderWidth: 1,
        backgroundColor: "#02458d",
      },
    ],
  },
  options: {
    scales: {
      y: {
        beginAtZero: true,
      },
    },
  },
});