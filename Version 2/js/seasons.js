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
// Обработчик событий на выбор даты , чтобы даты "От" и "До" не противоречили друг-другу
const data1 = document.getElementById("start-date");
const data2 = document.getElementById("end-date");
data1.addEventListener("input", () => {
  data2.min = data1.value;
});
data2.addEventListener("input", () => {
  data1.max = data2.value;
});
// Обработчик событий на первоначальный выбор направления,а после номер рейса
const direction = document.getElementById("Direction");
const flightSelection = document.getElementById("flightSelection");
direction.addEventListener("change", () => {
  flightSelection.readOnly = false;
  flightSelection.value = "";
});

// // Очищаем предыдущие элементы option, если они были
// flightSelection.innerHTML = "";

// let data = {
//   Direction: direction.value,
//   Date: date1.value,
// };
// // Number:
// fetch("/get_class", {
//   headers: {
// 	 Accept: "application/json",
// 	 "Content-Type": "application/json",
//   },
//   method: "POST",
//   body: JSON.stringify(data),
// })
//   .then((response) => {
// 	 return response.json();
//   })
//   .then((data) => {
// 	 data.rows.forEach((optionValue) => {
// 		const optionElement = document.createElement("option");
// 		optionElement.value = optionValue;
// 		optionElement.text = optionValue;
// 		flightSelection.appendChild(optionElement);
// 	 });
//   })
//   .catch((error) => {
// 	 console.log(error);
//   });
