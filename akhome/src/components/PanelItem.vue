<template>
  <div class="Box d-flex flex-column align-items-center" @click="Open" v-if="!smallIcon">
    <div class="Image-Board d-flex justify-content-center align-items-center">
      <img class="Image" :src="realImg" />
    </div>
    <div>{{ title }}</div>
  </div>

  <div class="Small" @click="Open" v-else>
    <div class="SmallImage-Board">
      <img class="SmallImage" :src="realImg" />
    </div>
    <div>{{ title }}</div>
  </div>
</template>

<script>
export default {
  name: "PanelItem",
  components: {},
  props: {
    img: String,
    title: String,
    url: String,
    replace: Boolean,
    smallIcon: Boolean,
  },
  data() {
    return {
      baseUrl: "",
      realImg: "",
    };
  },
  data() {
    return {
      baseUrl: "",
      realImg: "",
    };
  },
  methods: {
    Open() {
      if (this.url) {
        if (this.replace) {
          console.log("replace");
          window.location.replace(this.url);
        } else {
          console.log("open");
          window.open(this.url);
        }
      }
    },
  },
  mounted() {
    if (this.img.startsWith("http")) {
      this.realImg = this.img;
    } else {
      let token = localStorage.getItem("token") || "";
      this.realImg = this.baseUrl + "/icon?name=" + this.img + "&token=" + token;
    }
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

.Small {
  display: flex;
  width: 200px;
  margin-left: 10px;
  padding: 2px 10px;
  cursor: pointer;
  border: 0px solid;
  border-radius: 8px;
  align-items: center;
}

.Small:hover {
  box-shadow: 0px 0px 25px 3px rgb(0 0 0 / 7%), 0 10px 10px -5px rgb(0 0 0 / 5%) !important;
}

.SmallImage-Board {
  margin-right: 6px;
}

.SmallImage {
  max-width: 24px;
  max-height: 24px;
  object-fit: cover;
}
</style>
