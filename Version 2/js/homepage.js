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
const data = {
  labels: dates,
  datasets: [
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
      legend: { display: false, position: "bottom" },
      title: {
        display: false,
        text: "Сезонность для прогнозов (рейсы Москва-Сочи)",
      },
    },
  },
};

const myChart = new Chart(document.getElementById("myChart"), config);
const chartBody = document.querySelector(".chart__body");
const totalLabels = myChart.data.labels.length; // typo was fixed here
if (totalLabels > 30) {
  const newWidth = 1100 + (totalLabels - 30) * 40;
  chartBody.style.width = `${newWidth}px`;
}

// Обработчик событий на присваивание даты рейса на период вывода графика
const date1 = document.getElementById("date");
const date2 = document.getElementById("end-date");
const date3 = document.getElementById("start-date");
const flightSelection = document.getElementById("Number");

date1.addEventListener("input", () => {
  date2.max = date1.value;
  date3.max = date1.value;
  date2.value = date1.value;
  //   Устfнавливаем значение date3 на дату ровно на один месяц раньше, чем выбранная дата в date1
  const oneMonthEarlier = new Date(date1.value);
  oneMonthEarlier.setMonth(oneMonthEarlier.getMonth() - 1);
  date3.value = oneMonthEarlier.toISOString().slice(0, 10);

  // Очищаем предыдущие элементы option, если они были
  flightSelection.innerHTML = "";

  // Получаем новые данные из базы данных (замени этот код соответствующим кодом для базы данных)
  const newData = ["Option 1", "Option 2", "Option 3"];

  // Создаем новые элементы option и добавляем их в элемент flightSelection
  newData.forEach((optionValue) => {
    const optionElement = document.createElement("option");
    optionElement.value = optionValue;
    optionElement.text = optionValue;
    flightSelection.appendChild(optionElement);
  });
});
