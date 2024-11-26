<script>
  import { GreetService } from "../bindings/changeme";
  import { Events } from "@wailsio/runtime";

  let name = "";
  let result = "Please enter your name below 👇";
  let time = "Listening for Time event...";

  const doGreet = () => {
    let localName = name;
    if (!localName) {
      localName = "anonymous";
    }
    GreetService.Greet(localName)
      .then((resultValue) => {
        result = resultValue;
      })
      .catch((err) => {
        console.log(err);
      });
  };

  // Events.On("time", (timeValue) => {
  //   time = timeValue.data;
  // });
  Events.On("blur", () => {
    window;
  });
</script>

<div class="container">
  <input type="text" class="searchbar" />
</div>

<style>
  /* Put your standard CSS here */
  .searchbar {
    width: 100%;
    height: 100%;
    background: transparent;
    color: white;
    font-size: x-large;
  }
  .container {
    height: 80px;
    width: 600px;
  }
</style>
