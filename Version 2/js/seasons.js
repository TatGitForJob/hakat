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
// Обработчик событий на присваивание даты рейса на период вывода графика
const date2 = document.getElementById("end-date");
const date3 = document.getElementById("start-date");

function updateChart() {
  let chartExists = Chart.getChart("myChart");
  if (chartExists) {
    chartExists.destroy();
  }
  const myChart = new Chart(document.getElementById("myChart"), config);
  const chartBody = document.querySelector(".chart__body");
  const totalLabels = myChart.data.labels.length;
  if (totalLabels > 1) {
    const newWidth = 1100 + (totalLabels - 30) * 40;
    chartBody.style.width = `${newWidth}px`;
  }
}

date2.addEventListener("change", () => {
  let startDate = date3.value;
  let endDate = date2.value;
  myChart.data.labels = getDates(startDate, endDate);
  myChart.update();
  updateChart();
});

date3.addEventListener("change", () => {
  let startDate = date3.value;
  let endDate = date2.value;
  myChart.data.labels = getDates(startDate, endDate);
  myChart.update();
  updateChart();
});
let startDate = date3.value;
let endDate = date2.value;

function getDates(startDate, endDate) {
  const dates = [];
  let currentDate = new Date(startDate);
  const endDateTime = new Date(endDate).getTime();

  while (currentDate.getTime() <= endDateTime) {
    let day = ("0" + currentDate.getDate()).slice(-2);
    let month = ("0" + (currentDate.getMonth() + 1)).slice(-2);
    let year = currentDate.getFullYear();
    let formattedDate = `${day}-${month}-${year}`;
    dates.push(formattedDate);
    currentDate.setDate(currentDate.getDate() + 1);
  }
  return dates;
}

// График

// config
const config = {
  type: "bar",
  data: {
    labels: getDates(startDate, endDate),
    datasets: [
      {
        type: "line",
        label: "Line Dataset",
        data: [0],
        backgroundColor: ["rgba(255,0,0, 0.5)"],
        borderColor: ["red"],
        borderWidth: 4,
        fill: true,
        tension: 0.4,
        borderJoinStyle: "bevel",
      },
    ],
  },
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

myChart = new Chart(document.getElementById("myChart"), config);
chartBody = document.querySelector(".chart__body");
totalLabels = myChart.data.labels.length; // typo was fixed here
if (totalLabels > 1) {
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

askButton = document.getElementById("ask-Button");
clas = document.getElementById("Class");
number = document.getElementById("Number");
output = document.getElementById("output");
direction = document.getElementById("Direction");
direction.addEventListener("click", () => {
  askButton.disabled = false;
});

$(document).ready(function () {
  $("#download-Button").click(function () {
    // Отправка запроса на сервер
    window.location.href = "/download_season";
  });
});

askButton.addEventListener("click", function () {
  let data = {
    Direction: direction.value,
    Class: clas.value,
    Number: number.value,
    StartDate: data1.value,
    EndDate: data2.value,
  };
  // Number:

  fetch("/get_season", {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    method: "POST",
    body: JSON.stringify(data),
  })
    .then((response) => {
      response.text().then(function (data) {
        //output.textContent = JSON.parse(data);
        // Обновление данных графика
        myChart.data.datasets[0].data = JSON.parse(data).IntArray;
        myChart.data.labels = JSON.parse(data).StringArray;

        // Отображение графика при нажатие на кнопку
        let visibility = document.getElementById("visibility");
        visibility.classList.remove("visibility");

        // Проверка есть ли данные для построения графика,если нет то график не выводиться
        const isAllZeros = myChart.data.datasets[0].data.every((e) => e === 0);
        if (isAllZeros === true) {
          visibility.classList.add("visibility");
          const message = document.getElementById("message");
          message.textContent = "Отсутствуют данные";
        }

        myChart.update();
      });
    })
    .catch((error) => {
      console.log(error);
    });
});
