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

// // График
// const ctx = document.getElementById("myChart");

// new Chart(ctx, {
//   data: {
//     labels: [1, 2, 3, 4, 5, 6,1, 2, 3, 4, 5, 6,1, 2, 3, 4, 5, 6],
//     datasets: [
//       {
//         type: "line",
//         label: "Line Dataset",
//         data: [9, 12, 2, 3, 1, 2],
//         backgroundColor: ["rgba(255,0,0, 0.5)"],
//         borderColor: ["red"],
//         borderWidth: 4,
//         fill: true,
//         tension: 0.4,
//         borderJoinStyle: "bevel",
//       },
//       {
//         type: "bar",
//         label: "График сезонного спроса по классам бронирования",
//         data: [12, 19, 3, 5, 2, 3],
//         borderWidth: 1,
//         backgroundColor: ["#02458d", "grey", "green"],
//       },
//     ],
//   },
//   options: {
//     //  scales: {
//     //    y: {
//     //      beginAtZero: true,
//     //    },
//     //  },
//     plugins: {
// 			legend: { display: true, position: "bottom" },
// 			title: {
// 			display: true,
// 			text: "Сезонность для прогнозов (рейсы Москва-Сочи)",
// 			},
//       zoom: {
//         zoom: {
//           wheel: {
//             enabled: true,
// 				mode:'y',
//           },
//         },
//       },
//     },
//   },
// });
// График
const dates = Array.from({ length: 365 }, (_, i) => `День ${i + 1}`);
// setup
const data = {
  labels: dates,
  datasets: [
    {
      type: "line",
      label: "Line Dataset",
      data: [
        9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1,
        2, 9, 12, 2, 3, 1, 2, 9, 12, 2, 3, 1, 2,
      ],
      backgroundColor: ["rgba(255,0,0, 0.5)"],
      borderColor: ["red"],
      borderWidth: 4,
      fill: true,
      tension: 0.4,
      borderJoinStyle: "bevel",
    },
    {
      label: "График динамики бронирования",
      data: [
        18, 12, 6, 9, 12, 3, 9, 18, 12, 6, 9, 12, 3, 9, 18, 12, 6, 9, 12, 3, 9,
        18, 12, 6, 9, 12, 3, 9,
      ],
      backgroundColor: "#02458d",
      borderColor: "#02458d",
      borderWidth: 1,
    },
  ],
};

// config
const config = {
  type: "bar",
  data,
  options: {
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
        text: "Сезонность для прогнозов (рейсы Москва-Сочи)",
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
