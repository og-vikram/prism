<script>
  let isVisible = false; // Controls visibility of the spotlight search
  let searchQuery = ""; // The search input
  let results = []; // Mock search results
  let suggestions = ["Profile", "Settings", "Help", "Logout", "Dashboard"]; // Example data

  // Filters the suggestions based on the query
  const updateResults = () => {
    results = suggestions.filter((item) =>
      item.toLowerCase().includes(searchQuery.toLowerCase()),
    );
  };
</script>

<div class="searchbar">
  <input
    id="spotlight-input"
    type="text"
    placeholder="What do you want to do?"
  />
</div>

<style>
  .searchbar {
    position: fixed;
    top: 0;
    left: 0;
    border-radius: 10px;
    width: 100%;
    height: 100%;
  }

  .searchbar input {
    font-size: large;
    width: 100%;
    box-sizing: border-box;
    height: 100%;
    background: none;
    color: white;
    border: none;
    padding-inline: 10px;
  }
</style>
