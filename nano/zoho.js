// zoho
console.log("hello zoho!");

// scene
const sce = {
	init() {
		sce.a = {};
		sce.a.r = [1920,1080,0,0];
		sce.a.c = dom.canvas(sce.a.r);
		dom.add(sce.a.c);
		console.log(sce.a);
	},
	scene() {
		// webgl
		sce.b = new THREE.Scene();
		// cam 
		sce.c = new THREE.OrthographicCamera(0,1920,0,1080,-1000,1000);
		// render
		sce.d = new THREE.WebGLRenderer({canvas:sce.a.c});
	},
	debug() {
		console.log(sce.d.info);
		console.log(sce.d.getContext());
	}
}


