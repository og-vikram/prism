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
</script>

<div class="container">
  <h1 class="greeting">Hello!</h1>
</div>

<style>
  /* Put your standard CSS here */
  h1 {
    color: white;
  }
  .container {
    background-color: blue;
  }
</style>
