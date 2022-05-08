<template>
  <div>
    <div class="Title">AlwaysKing HomeLab</div>

    <div class="d-flex flex-wrap justify-content-center">
      <PanelGroup
        v-for="Group in Apps"
        :key="Group.Title"
        :Title="Group.Title"
        :Image="Group.Image"
      >
        <div
          class="
            Panel
            d-flex
            align-content-start
            justify-content-start
            flex-wrap
          "
        >
          <PanelItem
            v-for="item in Group.List"
            :key="item"
            :img="item.Image"
            :title="item.AppName"
            :url="item.URL"
            :uri="item.URI"
          ></PanelItem>
        </div>
      </PanelGroup>
    </div>
  </div>
</template>

<script>
import PanelItem from "./components/PanelItem.vue";
import PanelGroup from "./components/PanelGroup.vue";

export default {
  name: "App",
  components: {
    PanelItem,
    PanelGroup,
  },
  data() {
    return {
      Apps: [],
    };
  },
  mounted() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/App.conf", false);
    xmlHttp.send(null);
    this.Apps = JSON.parse(xmlHttp.responseText);
  },
};
</script>

<style>
#app {
  font-family: Roboto, Helvetica, Arial, sans-serif;
  height: 100%;
  width: 100%;
}

body,
html {
  height: 100%;
  width: 100%;
  background-color: rgb(251, 251, 251);
}
.Title {
  height: 100px;
  font-size: 30px;
  text-align: center;
}
.Panel {
  height: 100%;
  width: 100%;
}
</style>
