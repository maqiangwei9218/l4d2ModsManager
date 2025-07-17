import axios from "axios";

function logExit() {
  // const data = {
  //   url: window.location.href,
  //   timestamp: Date.now(),
  //   userAgent: navigator.userAgent,
  // };

  // if (navigator.sendBeacon) {
  //   // 使用 sendBeacon
  //   navigator.sendBeacon("/api/log-exit", JSON.stringify(data));
  // } else {
  //   // 回退方案 - 同步 XHR
  //   try {
  //     var xhr = new XMLHttpRequest();
  //     xhr.open("POST", "/api/log-exit", false);
  //     xhr.setRequestHeader("Content-Type", "application/json");
  //     xhr.send(JSON.stringify(data));
  //   } catch (e) {
  //     // 错误处理
  //   }
  // }

  axios("http://127.0.0.1:9090/exit");
}

// 同时监听多个事件以确保覆盖所有情况
window.addEventListener("pagehide", logExit);
window.addEventListener("beforeunload", logExit);
window.addEventListener("unload", logExit);
