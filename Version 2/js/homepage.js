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

askButton = document.getElementById("ask-Button");
clas = document.getElementById("Class");
number = document.getElementById("Number");
output = document.getElementById("output");

$(document).ready(function () {
  $("#download-Button").click(function () {
    // Отправка запроса на сервер
    window.location.href = "/download_dinamic";
  });
});

askButton.addEventListener("click", function () {
  let data = {
    Direction: direction.value,
    Date: date1.value,
    Class: clas.value,
    Number: number.value,
    StartDate: date2.value,
    EndDate: date3.value,
  };
  // Number:

  fetch("/get_dinamic", {
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
        myChart.data.datasets[0].data = JSON.parse(data);

        // Отображение графика при нажатие на кнопку
        let visibility = document.getElementById("visibility");
        visibility.classList.remove("visibility");

        // Проверка есть ли данные для построения графика,если нет то график не выводиться
        const isAllZeros = myChart.data.datasets[0].data.every((e) => e === 0);
        if (isAllZeros === true) {
          visibility.classList.add("visibility");
          const message = document.getElementById("message");
          message.textContent = "Отсутствуют данные";
        } else {
          const message = document.getElementById("message");
          message.textContent = "";
        }
        const isAllEmpty = myChart.data.datasets[0].data;
        if (isAllEmpty.length === 0) {
          visibility.classList.add("visibility");
          const message = document.getElementById("message");
          message.textContent = "Пустой массив";
        } else {
          const message = document.getElementById("message");
          message.textContent = "";
        }

        myChart.update();
      });
    })
    .catch((error) => {
      console.log(error);
    });
  //  Добавляем в график период для просмотра динамики бронирования
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
  getDates(startDate, endDate);
  const dates = getDates(startDate, endDate);
  myChart.data.labels = dates;
  myChart.update();
});

// График
const dates1 = Array.from({ length: 150 }, (_, i) => `День ${i + 1}`);

const config = {
  type: "bar",
  data: {
    labels: dates1,
    datasets: [
      {
        label: "График динамики бронирования",
        data: [],
        backgroundColor: "#02458d",
        borderColor: "#02458d",
        borderWidth: 1,
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
      legend: { display: false, position: "bottom" },
      title: {
        display: false,
        text: "Сезонность для прогнозов (рейсы Москва-Сочи)",
      },
    },
  },
};

// Обработчик событий на присваивание даты рейса на период вывода графика
const date1 = document.getElementById("date");
const date2 = document.getElementById("end-date");
const date3 = document.getElementById("start-date");
let direction = document.getElementById("Direction");
const flightSelection = document.getElementById("Number");

date1.addEventListener("input", () => {
  date2.max = date1.value;
  date3.max = date1.value;
  date2.value = date1.value;
  askButton.disabled = false;
  const oneMonthEarlier = new Date(date1.value);
  oneMonthEarlier.setMonth(oneMonthEarlier.getMonth() - 1);
  date3.value = oneMonthEarlier.toISOString().slice(0, 10);

  // Очищаем предыдущие элементы option, если они были
  flightSelection.innerHTML = "";

  let data = {
    Direction: direction.value,
    Date: date1.value,
  };
  // Number:
  fetch("/get_class", {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    method: "POST",
    body: JSON.stringify(data),
  })
    .then((response) => {
      return response.json();
    })
    .then((data) => {
      data.rows.forEach((optionValue) => {
        const optionElement = document.createElement("option");
        optionElement.value = optionValue;
        optionElement.text = optionValue;
        flightSelection.appendChild(optionElement);
      });
    })
    .catch((error) => {
      console.log(error);
    });
});
myChart = new Chart(document.getElementById("myChart"), config);
chartBody = document.querySelector(".chart__body");
totalLabels = myChart.data.labels.length;
if (totalLabels > 1) {
  const newWidth = 1100 + (totalLabels - 30) * 40;
  chartBody.style.width = `${newWidth}px`;
}
