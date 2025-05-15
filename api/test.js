const DefaultApi = require('./js-client/dist/api/DefaultApi.js').default;
const ApiClient = require('./js-client/dist/ApiClient.js').default;
const UserNameActionRequest = require('./js-client/dist/model/UserNameActionRequest.js').default;
const UserName = require('./js-client/dist/model/UserName.js').default;

const client = ApiClient.instance || new ApiClient();
client.basePath = "http://localhost:8080";

const api = new DefaultApi(client);

// 构造请求体
const user = new UserName('张三');
const req = new UserNameActionRequest('get', user);

async function main() {
    try {
        const result = await api.userActions(req);
        console.log('返回结果：', result);
    } catch (err) {
        console.error('调用出错：', err);
    }
}

main();