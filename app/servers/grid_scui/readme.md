# 设计规则
view    试图, 需要又vue文件，设计插槽
Render  组件，可以任何东西，放入到view的插槽

# 固定插槽
("/>)不支持有空格

    <slot id="Render的名称"/>

# 全局对象
任意地方都可以调用的对象代码

    $API {admin_user: undefined, auth: {…}, common: {…}, demo: {…}, load: undefined, …}
    $AUTH ƒ permission(data)
    $CONFIG {APP_NAME: 'SCUI(DEV)', DASHBOARD_URL: '/dashboard', APP_VER: '1.6.6', CORE_VER: '1.6.6', API_URL: 'http://127.0.0.1:8080/admin', …}
    $HTTP {get: ƒ, post: ƒ, put: ƒ, patch: ƒ, delete: ƒ, …}
    $ROLE ƒ rolePermission(data)
    $TOOL {data: {…}, session: {…}, cookie: {…}, screen: ƒ, objCopy: ƒ, …}
    $alert (message, titleOrOpts, options, appContext) => {…}
    $confirm (message, titleOrOpts, options, appContext) => {…}
    $loading ƒ (options = {})
    $message (options = {}, context) => {…}
    $messageBox  MessageBox(options, appContext = null)
    $msgbox  MessageBox(options, appContext = null)
    $notify ƒ (options = {}, context = null)
    $prompt
    $route
    $router
    $store
