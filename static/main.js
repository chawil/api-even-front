window.onload = async function () {
  var numberInputElement = document.querySelector("#number");
  var isEvenDivElement = document.querySelector("#isEven");

  const checkNumber = async (number) => {
    var response = await (await fetch("/api/even", { method: "POST", body: number })).json();
    isEvenDivElement.innerHTML = JSON.stringify(response);
  };
  numberInputElement.onchange = (e) => checkNumber(e.target.value);
  checkNumber(numberInputElement.value);
};
