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

  fetch("/get_time", {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    method: "POST",
    body: JSON.stringify(data),
  })
    .then((response) => {
      response.text().then(function (data) {
        output.textContent = JSON.parse(data);
        let array = JSON.parse(data);
        // Обновление данных графика
        myChart.data.datasets[0].data = array;
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
        data: [
          1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3,
          4, 5,
        ],
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
const direction = document.getElementById("Direction");
const flightSelection = document.getElementById("Number");

date1.addEventListener("input", () => {
  date2.max = date1.value;
  date3.max = date1.value;
  date2.value = date1.value;
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
const myChart = new Chart(document.getElementById("myChart"), config);
const chartBody = document.querySelector(".chart__body");
const totalLabels = myChart.data.labels.length; // typo was fixed here
if (totalLabels > 30) {
  const newWidth = 1100 + (totalLabels - 30) * 40;
  chartBody.style.width = `${newWidth}px`;
}
