const DefaultApi = require('./js-client/dist/api/DefaultApi.js').default;
const ApiClient = require('./js-client/dist/ApiClient.js').default;
const UserNameActionRequest = require('./js-client/dist/model/UserNameActionRequest.js').default;
const UserName = require('./js-client/dist/model/UserName.js').default;
const UserInfo = require('./js-client/dist/model/UserInfo.js').default;

const client = ApiClient.instance || new ApiClient();
client.basePath = "http://localhost:8080";

const api = new DefaultApi(client);

// 构造请求体
const user = new UserName('张三');
const req = new UserNameActionRequest('get', user);

api.getUser(req, (error, data, response) => {
    if (error) {
      console.error('调用出错：', error);
    } else {
      console.log('返回结果：', data);
      // 解析 user-json 字段
      if (data && data['user-json']) {
        try {
          const userJsonObj = UserInfo.constructFromObject(JSON.parse(data['user-json']));
          console.log('解析后的 user-json：', userJsonObj);
        } catch (e) {
          console.error('user-json 解析失败：', e);
        }
      }
    }
  });
