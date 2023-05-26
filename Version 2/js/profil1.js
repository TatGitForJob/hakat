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
const dates = Array.from({ length: 365 }, (_, i) => `День ${i + 1}`);
// setup
const data = {
  labels: dates,
  datasets: [
    //  График Спроса А
    {
      type: "line",
      label: "Спрос А",
      data: [
        9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1,
        2, 9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1, 2,
      ],
      pointRadius: 0,
      order: 5,
    },
    //  График Спроса Б
    {
      type: "line",
      label: "Спрос Б",
      data: [
        18, 12, 6, 9, 12, 3, 9, 18, 12, 6, 9, 12, 3, 9, 18, 12, 6, 9, 12, 3, 9,
        18, 12, 6, 9, 12, 3, 9,
      ],
      backgroundColor: ["rgba(123, 121, 209, 1)"],
      borderColor: ["rgba(123, 121, 209, 1)"],
      borderWidth: 4,
      fill: false,
      tension: 0.1,
      borderJoinStyle: "bevel",
      order: 1,
      pointRadius: 0,
    },
    //  График Ожидаемое бронирование Б
    {
      type: "line",
      label: "Ожидаемое бронирование Б",
      data: [
        6, 9, 12, 3, 9, 18, 12, 6, 9, 12, 3, 9, 18, 12, 6, 9, 12, 3, 9, 18, 12,
        6, 9, 12, 3, 9,
      ],
      backgroundColor: ["rgba(255, 0, 0, 1)"],
      borderColor: ["rgba(255, 0, 0, 1)"],
      borderWidth: 4,
      fill: false,
      tension: 0.1,
      borderJoinStyle: "bevel",
      order: 1,
      pointRadius: 0,
    },
    //  График Ожидаемое бронирование AБ
    {
      type: "scatter",
      label: "Ожидаемое бронирование AБ",
      data: [3, 4, 4, 3, 4, 23, 4, 2, 34],
      backgroundColor: ["rgba(133, 127, 127, 1)"],
      borderColor: ["rgba(133, 127, 127, 1)"],
      borderWidth: 4,
      fill: false,
      tension: 0.1,
      borderJoinStyle: "bevel",
      order: 1,
    },
  ],
};

// config
const config = {
  type: "line",
  data,

  options: {
    backgroundColor: ["rgba(41, 39, 156, 1)"],
    borderColor: ["rgba(41, 39, 156, 1)"],
    borderWidth: 4,
    fill: true,
    tension: 0.4,
    borderJoinStyle: "bevel",
    maintainAspectRatio: false,

    scales: {
      y: {
        beginAtZero: true,
      },
    },
    plugins: {
      zoom: {
        zoom: {
          wheel: {
            enabled: true,
            mode: "y",
            modifierKey: "ctrl",
          },
        },
      },
      legend: {
        display: false,
        position: "left",
      },
      title: {
        display: false,
        text: "Спрос А)",
      },
    },
  },
};

// render init block
const myChart = new Chart(document.getElementById("myChart"), config);
const chartBody = document.querySelector(".chart__body");
const totalLabels = myChart.data.labels.length; // typo was fixed here
if (totalLabels > 30) {
  const newWidth = 1100 + (totalLabels - 30) * 40;
  chartBody.style.width = `${newWidth}px`;
}
