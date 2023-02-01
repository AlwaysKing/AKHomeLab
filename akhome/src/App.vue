<template>
  <div>
    <div class="Title">AlwaysKing HomeLab</div>

    <div v-if="!Login">
      <div class="SubTitle">
        常用
        <div class="SwitchBtn" @click="CheckTopView">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" v-if="TopSmall">
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M2 8C2 6.34315 3.34315 5 5 5H19C20.6569 5 22 6.34315 22 8V16C22 17.6569 20.6569 19 19 19H5C3.34315 19 2 17.6569 2 16V8ZM16 7H19C19.5523 7 20 7.44771 20 8V16C20 16.5523 19.5523 17 19 17H16V7ZM14 7H10V17H14V7ZM8 17V7H5C4.44772 7 4 7.44772 4 8V16C4 16.5523 4.44772 17 5 17H8Z"
              fill="currentColor"
            />
          </svg>

          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" v-else>
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M2 8C2 6.34315 3.34315 5 5 5H19C20.6569 5 22 6.34315 22 8V16C22 17.6569 20.6569 19 19 19H5C3.34315 19 2 17.6569 2 16V8ZM13 7H19C19.5523 7 20 7.44771 20 8V16C20 16.5523 19.5523 17 19 17H13V7ZM11 7H5C4.44772 7 4 7.44772 4 8V16C4 16.5523 4.44772 17 5 17H11V7Z"
              fill="currentColor"
            />
          </svg>
        </div>
      </div>
      <div class="d-flex flex-wrap justify-content-center">
        <div class="Panel d-flex align-content-start justify-content-center flex-wrap">
          <PanelItem
            v-for="item in Recent"
            :key="item"
            :img="item.Image"
            :title="item.AppName"
            :url="item.URL"
            :uri="item.URI"
            :SmallIcon="TopSmall"
          ></PanelItem>
        </div>
      </div>
      <div class="SubTitle" style="margin-top: 30px">
        全部
        <div class="SwitchBtn" @click="CheckAppsView">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" v-if="AppsSmall">
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M2 8C2 6.34315 3.34315 5 5 5H19C20.6569 5 22 6.34315 22 8V16C22 17.6569 20.6569 19 19 19H5C3.34315 19 2 17.6569 2 16V8ZM16 7H19C19.5523 7 20 7.44771 20 8V16C20 16.5523 19.5523 17 19 17H16V7ZM14 7H10V17H14V7ZM8 17V7H5C4.44772 7 4 7.44772 4 8V16C4 16.5523 4.44772 17 5 17H8Z"
              fill="currentColor"
            />
          </svg>

          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" v-else>
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M2 8C2 6.34315 3.34315 5 5 5H19C20.6569 5 22 6.34315 22 8V16C22 17.6569 20.6569 19 19 19H5C3.34315 19 2 17.6569 2 16V8ZM13 7H19C19.5523 7 20 7.44771 20 8V16C20 16.5523 19.5523 17 19 17H13V7ZM11 7H5C4.44772 7 4 7.44772 4 8V16C4 16.5523 4.44772 17 5 17H11V7Z"
              fill="currentColor"
            />
          </svg>
        </div>
      </div>
      <div class="d-flex flex-wrap justify-content-center">
        <PanelGroup v-for="Group in Apps" :key="Group.Title" :Title="Group.Title" :Image="Group.Image" :SmallIcon="AppsSmall">
          <div class="Panel d-flex align-content-start justify-content-start flex-wrap" :class="{ AppSmallIcon: AppsSmall }">
            <PanelItem
              v-for="item in Group.List"
              :key="item"
              :img="item.Image"
              :title="item.AppName"
              :url="item.URL"
              :uri="item.URI"
              :SmallIcon="AppsSmall"
            ></PanelItem>
          </div>
        </PanelGroup>
      </div>
    </div>
    <div class="d-flex justify-content-center" v-else>
      <div style="width: 220px">
        <div class="d-flex">
          <MDBInput inputGroup :formOutline="false" wrapperClass="mb-3" v-model="User" placeholder="用户名">
            <MDBBtn outline="primary" @click="SendAuth"> 发送 </MDBBtn>
          </MDBInput>
          <div v-if="UserError == 1" style="color: green" class="Status">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M10.2426 16.3137L6 12.071L7.41421 10.6568L10.2426 13.4853L15.8995 7.8284L17.3137 9.24262L10.2426 16.3137Z" fill="currentColor" />
              <path
                fill-rule="evenodd"
                clip-rule="evenodd"
                d="M1 12C1 5.92487 5.92487 1 12 1C18.0751 1 23 5.92487 23 12C23 18.0751 18.0751 23 12 23C5.92487 23 1 18.0751 1 12ZM12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21Z"
                fill="currentColor"
              />
            </svg>
          </div>

          <div v-if="UserError == 2" style="color: red" class="Status">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path
                d="M16.3394 9.32245C16.7434 8.94589 16.7657 8.31312 16.3891 7.90911C16.0126 7.50509 15.3798 7.48283 14.9758 7.85938L12.0497 10.5866L9.32245 7.66048C8.94589 7.25647 8.31312 7.23421 7.90911 7.61076C7.50509 7.98731 7.48283 8.62008 7.85938 9.0241L10.5866 11.9502L7.66048 14.6775C7.25647 15.054 7.23421 15.6868 7.61076 16.0908C7.98731 16.4948 8.62008 16.5171 9.0241 16.1405L11.9502 13.4133L14.6775 16.3394C15.054 16.7434 15.6868 16.7657 16.0908 16.3891C16.4948 16.0126 16.5171 15.3798 16.1405 14.9758L13.4133 12.0497L16.3394 9.32245Z"
                fill="currentColor"
              />
              <path
                fill-rule="evenodd"
                clip-rule="evenodd"
                d="M1 12C1 5.92487 5.92487 1 12 1C18.0751 1 23 5.92487 23 12C23 18.0751 18.0751 23 12 23C5.92487 23 1 18.0751 1 12ZM12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21Z"
                fill="currentColor"
              />
            </svg>
          </div>
        </div>
        <div class="d-flex">
          <MDBInput inputGroup :formOutline="false" wrapperClass="mb-3" v-model="Auth" placeholder="验证码">
            <MDBBtn outline="primary" @click="GenToken"> 验证 </MDBBtn>
          </MDBInput>
          <div v-if="LoginError != 0" style="color: red" class="Status">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path
                d="M16.3394 9.32245C16.7434 8.94589 16.7657 8.31312 16.3891 7.90911C16.0126 7.50509 15.3798 7.48283 14.9758 7.85938L12.0497 10.5866L9.32245 7.66048C8.94589 7.25647 8.31312 7.23421 7.90911 7.61076C7.50509 7.98731 7.48283 8.62008 7.85938 9.0241L10.5866 11.9502L7.66048 14.6775C7.25647 15.054 7.23421 15.6868 7.61076 16.0908C7.98731 16.4948 8.62008 16.5171 9.0241 16.1405L11.9502 13.4133L14.6775 16.3394C15.054 16.7434 15.6868 16.7657 16.0908 16.3891C16.4948 16.0126 16.5171 15.3798 16.1405 14.9758L13.4133 12.0497L16.3394 9.32245Z"
                fill="currentColor"
              />
              <path
                fill-rule="evenodd"
                clip-rule="evenodd"
                d="M1 12C1 5.92487 5.92487 1 12 1C18.0751 1 23 5.92487 23 12C23 18.0751 18.0751 23 12 23C5.92487 23 1 18.0751 1 12ZM12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21Z"
                fill="currentColor"
              />
            </svg>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import PanelItem from "./components/PanelItem.vue";
import PanelGroup from "./components/PanelGroup.vue";
import { MDBInput, MDBBtn } from "mdb-vue-ui-kit";
import { MDBIcon } from "mdb-vue-ui-kit";

export default {
  name: "App",
  components: {
    PanelItem,
    PanelGroup,
    MDBInput,
    MDBBtn,
    MDBIcon,
  },
  data() {
    return {
      url: "",
      Login: false,
      User: "",
      Auth: "",
      UserError: 0,
      LoginError: 0,
      TopSmall: false,
      AppsSmall: true,
      Apps: [],
      Recent: [],
    };
  },
  methods: {
    CheckAppsView() {
      this.AppsSmall = !this.AppsSmall;
      localStorage.setItem("AppsSmall", this.AppsSmall ? "true" : "false");
    },
    CheckTopView() {
      this.TopSmall = !this.TopSmall;
      localStorage.setItem("TopSmall", this.TopSmall ? "true" : "false");
    },
    GetData() {
      let token = localStorage.getItem("token") || "";
      var xmlHttp = new XMLHttpRequest();
      xmlHttp.open("GET", this.url + "/data", false);
      xmlHttp.setRequestHeader("token", token);
      xmlHttp.send(null);
      var data = JSON.parse(xmlHttp.responseText);
      if (data.status != false) {
        this.Apps = data.Apps;
        var AllApps = {};
        this.Apps.forEach((Group) => {
          Group.List.forEach((item) => {
            AllApps[item.AppName] = item;
          });
        });

        data.Recent.forEach((element) => {
          this.Recent.push(AllApps[element]);
        });
      } else {
        this.Login = true;
      }
    },
    //
    RefreshToken() {
      let token = localStorage.getItem("token") || "";
      var xmlHttp = new XMLHttpRequest();
      xmlHttp.open("GET", this.url + "/refreshtoken", false);
      xmlHttp.setRequestHeader("token", token);
      xmlHttp.send(null);
      var data = JSON.parse(xmlHttp.responseText);
      if (data.status) {
        localStorage.setItem("token", data.token);
      }
    },
    CheckToken() {
      let token = localStorage.getItem("token") || "";
      if (token.length == 0) {
        return false;
      } else {
        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open("GET", this.url + "/checktoken", false);
        xmlHttp.setRequestHeader("token", token);
        xmlHttp.send(null);
        var data = JSON.parse(xmlHttp.responseText);
        if (data.vaild == true && data.refresh == true) {
          this.RefreshToken();
        }
        return data.vaild;
      }
    },
    SendAuth() {
      this.UserError = 0;
      let token = localStorage.getItem("token") || "";
      var xmlHttp = new XMLHttpRequest();
      xmlHttp.open("GET", this.url + "/sendauth?name=" + this.User, false);
      xmlHttp.setRequestHeader("token", token);
      xmlHttp.send(null);
      var data = JSON.parse(xmlHttp.responseText);
      if (data.status != true) {
        this.UserError = 2;
      } else {
        this.UserError = 1;
      }
    },
    GenToken() {
      this.LoginError = 0;
      let token = localStorage.getItem("token") || "";
      var xmlHttp = new XMLHttpRequest();
      xmlHttp.open("GET", this.url + "/gentoken?name=" + this.User + "&auth=" + this.Auth, false);
      xmlHttp.setRequestHeader("token", token);
      xmlHttp.send(null);
      var data = JSON.parse(xmlHttp.responseText);
      if (data.status != true) {
        this.LoginError = 1;
        return false;
      } else {
        localStorage.setItem("token", data.token);
        this.Login = false;
        this.GetData();
        return true;
      }
    },
  },
  mounted() {
    this.AppsSmall = localStorage.getItem("AppsSmall") == "true";
    this.TopSmall = localStorage.getItem("TopSmall") == "true";

    if (this.CheckToken()) {
      this.GetData();
    } else {
      this.Login = true;
    }
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
  /* text-align: center; */
  display: flex;
  justify-content: center;
}

.SubTitle {
  display: flex;
  justify-content: center;
  font-size: 25px;
}

.Panel {
  height: 100%;
  width: 100%;
  max-width: 1400px;
}

.AppSmallIcon {
  flex-direction: column;
}

.SwitchBtn {
  margin-left: 12px;
  cursor: pointer;
}
.Status {
  margin-left: 230px;
  margin-top: 4px;
  position: absolute;
}
</style>
