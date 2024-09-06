<script setup lang="ts">
import { reactive } from "vue";
import UsFlag from "./assets/icon/flag/Us.vue";
import CheckmarkCircle from "./assets/icon/CheckmarkCircle.vue";
import { NodeInfo } from "./models/node.ts";

const nodeMaps = new Map<string, NodeInfo>();

const nodeData = reactive({
  nodeMaps: nodeMaps,
});

const socket = new WebSocket('ws://localhost:8080/web');

socket.onopen = () => {
  console.log('WebSocket connection established.');
};

socket.onmessage = (event) => {
  const data = event.data;

  const nodeInfo: NodeInfo = JSON.parse(data);
  // nodeMaps.set(nodeInfo.nodeId, nodeInfo);

  // 使用 Vue.set 或者直接赋值给响应式对象的属性来确保响应式更新
  nodeData.nodeMaps.set(nodeInfo.nodeId, nodeInfo);
  console.log(nodeInfo.nodeId);

  console.log('Received data:', data);


  let nodeInfos: NodeInfo[] = Array.from(nodeMaps.values());

  console.log("nodeInfos", nodeInfos);

};

socket.onerror = (error) => {
  console.error('WebSocket error:', error);
};

socket.onclose = (event) => {
  if (event.wasClean) {
    console.log('Connection closed cleanly.');
  } else {
    console.log('Connection closed unexpectedly.');
  }
};


// 计算属性用于格式化网络速度
const formattedNetworkSpeed = (speed: number) => {
  const speedKB = speed / 1024;
  if (speedKB > 1024) {
    return `${ (speedKB / 1024).toFixed(2) } MB/s`;
  } else {
    return `${ speedKB.toFixed(2) } KB/s`;
  }
};
</script>

<template>
  <n-grid cols="1">
    <n-gi>
      <div style="height: 8vh;background-color: #eff0f5;display: flex;align-items: center;">
        <div style="display: flex;align-items: center;justify-content: center;margin-left: 2rem">
          <n-icon :component="CheckmarkCircle" size="40" color="lightblue"></n-icon>
          <h1>LiveProbe 管理系统</h1>
        </div>
      </div>
    </n-gi>
    <n-gi>
      <div style="display: flex;width: 100vw;height: 84vh;flex-wrap: wrap;overflow: auto;">
        <n-card hoverable class="node-card" v-for="node in Array.from(nodeData.nodeMaps.values())">
          <template #header>
            <div style="display: flex;">
              <div style="display: flex;align-items: center;margin-right: 0.5rem ">
                <n-icon :component="UsFlag" size="30"></n-icon>
              </div>
              {{ node.nodeId.substring(0, 8) }}
            </div>
          </template>

          <template #header-extra>
            <n-tag type="success">
              <template #icon>
                <n-icon :component="CheckmarkCircle"/>
              </template>
              {{ node.host.bootDays }}天
            </n-tag>
          </template>
          <n-grid cols="1" y-gap="10px">
            <n-gi>
              <n-grid cols="2">

<!--                <n-gi>-->
<!--                  <div style="text-align: center">负载</div>-->
<!--                  <div style="text-align: center">-->
<!--                    <n-progress style="width: 80%;" type="dashboard" gap-position="bottom"-->
<!--                                :percentage="(node.cpu.usedPercent * (3 / 5) + node.memory.usedPercent * (2 / 5)).toFixed(2)">-->

<!--                      <span style="text-align: center">{{-->
<!--                          (node.cpu.usedPercent * (3 / 4) + node.memory.usedPercent * (1 / 4)).toFixed(2)-->
<!--                        }}% </span>-->
<!--                    </n-progress>-->
<!--                  </div>-->
<!--                </n-gi>-->

                <n-gi>
                  <div style="text-align: center">CPU</div>

                  <n-tooltip trigger="hover">
                    <template #trigger>
                      <div style="text-align: center">
                        <n-progress style="width: 80%;" type="dashboard" gap-position="bottom"
                                    :percentage="node.cpu.usedPercent">
                          <span style="text-align: center">{{ (node.cpu.usedPercent).toFixed(2) }}%</span>
                        </n-progress>
                      </div>
                    </template>

                    <p>型号：{{ node.cpu.modelName }}</p>
                    <p>核心：{{ node.cpu.core }}</p>
                    <p>主频：{{ node.cpu.mhz }} MHz</p>
                    <p>缓存：{{ node.cpu.cacheSize }}</p>

                  </n-tooltip>

                </n-gi>


                <n-gi>

                  <div style="text-align: center">内存</div>
                  <n-tooltip trigger="hover">
                    <template #trigger>
                      <div style="text-align: center">
                        <n-progress style="width: 80%;" type="dashboard" gap-position="bottom"
                                    :percentage="(node.memory.usedPercent).toFixed(2)">

                          <span style="text-align: center">{{ (node.memory.usedPercent).toFixed(2) }}% </span>
                        </n-progress>
                      </div>

                    </template>

                    <p>总量：{{ (node.memory.total / 1024).toFixed(2) }} G</p>
                    <p>使用：{{ (node.memory.used / 1024).toFixed(2) }} G</p>
                    <p>空闲：{{ (node.memory.available / 1024).toFixed(2) }} G</p>

                  </n-tooltip>

                </n-gi>
                <n-gi>
                  <div style="text-align: center">硬盘</div>
                  <n-tooltip trigger="hover">
                    <template #trigger>
                      <div style="text-align: center">
                        <n-progress style="width: 80%;" type="dashboard" gap-position="bottom"
                                    :percentage="node.disk.usedPercent.toFixed(2)">
                          <span style="text-align: center">{{ (node.disk.usedPercent).toFixed(2) }}% </span>

                        </n-progress>
                      </div>
                    </template>

                    <p>总量：{{ (node.disk.total / 1024).toFixed(2) }} G</p>
                    <p>使用：{{ (node.disk.used / 1024).toFixed(2) }} G</p>
                    <p>空闲：{{ (node.disk.available / 1024).toFixed(2) }} G</p>
                    <p>挂载：{{ node.disk.mountPath }}</p>

                  </n-tooltip>

                </n-gi>

                <n-gi>
                  <div style="text-align: center">SWAP</div>
                  <n-tooltip trigger="hover">
                    <template #trigger>
                      <div style="text-align: center">
                        <n-progress style="width: 80%;" type="dashboard" gap-position="bottom"
                                    :percentage="(node.swap.usedPercent).toFixed(2)">
                          <span style="text-align: center">{{ (node.swap.usedPercent).toFixed(2) }}% </span>

                        </n-progress>
                      </div>
                    </template>

                    <p>总量：{{ (node.swap.total / 1024 / 1024).toFixed(2) }} MB</p>
                    <p>使用：{{ (node.swap.used / 1024/ 1024).toFixed(2) }} MB</p>
                    <p>空闲：{{ (node.swap.available / 1024/ 1024).toFixed(2) }} MB</p>

                  </n-tooltip>

                </n-gi>
              </n-grid>
            </n-gi>
            <n-gi>
              网络
              <n-tag type="info">
                ↑ {{ formattedNetworkSpeed(node.network.uploadSpeed) }}
              </n-tag>
              <n-tag type="warning" style="margin-left: 8px">
                ↓ {{ formattedNetworkSpeed(node.network.downloadSpeed) }}
              </n-tag>
            </n-gi>



          </n-grid>
        </n-card>
      </div>
    </n-gi>

    <n-gi>
      <div style="height: 8vh;background-color: #eff0f5;display: flex;align-items: center;justify-content: center;">
        <h1>这是 Footer </h1>
      </div>
    </n-gi>


  </n-grid>


</template>


<style scoped>
.node-card {
  margin: 10px;
  max-width: 320px;
  height: 450px;
  border-radius: .75rem;
  background-color: #f1f2f5;
}

.node-card:hover {
  background-color: #ffffff;
  border: 1px solid #e4e7ed;
}

</style>
