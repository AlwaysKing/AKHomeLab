<template>
  <div class="Box d-flex flex-column align-items-center" @click="Open">
    <div class="Image-Board d-flex justify-content-center align-items-center">
      <img class="Image" :src="img" />
    </div>

    <div>{{ title }}</div>
  </div>
</template>

<script>
// import { MDBContainer } from "mdb-vue-ui-kit";

export default {
  name: "PanelItem",
  components: {},
  props: {
    img: String,
    title: String,
    url: String,
    uri: String,
  },
  methods: {
    Open() {
      if (this.url) {
        window.open(this.url + this.uri);
      }
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.Box {
  width: 180px;
  height: 230px;
  margin: 20px;
  border: 0px solid;
  border-radius: 8px;
  box-shadow: 0 10px 15px -3px rgb(0 0 0 / 7%), 0 4px 6px -2px rgb(0 0 0 / 5%) !important;
  cursor: pointer;
}

.Box:hover {
  box-shadow: 0px 0px 25px 3px rgb(0 0 0 / 7%), 0 10px 10px -5px rgb(0 0 0 / 5%) !important;
}
.Image-Board {
  width: 180px;
  height: 180px;
}
.Image {
  max-width: 150px;
  max-height: 160px;
  object-fit: cover;
}
</style>
