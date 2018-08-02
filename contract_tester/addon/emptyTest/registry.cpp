#include <node.h>
void RegisterModule(v8::Handle<v8::Object> target) {
		// 注册模块功能，负责导出接口到node.js
}
// 注册模块名称，编译后，模块将编译成modulename.node文件
// 当你需要修改模块名字的时候，需要修改 binding.gyp("target_name") 和此处
NODE_MODULE(registry, RegisterModule);
