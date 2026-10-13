console.log("🚀 app.js is starting...");

const go = new Go();
WebAssembly.instantiateStreaming(fetch("solver.wasm"), go.importObject).then((result) => {
    go.run(result.instance); 
    
    document.getElementById("status").innerText = "Engine Ready 🚀";
    document.getElementById("status").style.color = "#4CAF50";
    document.getElementById("solveBtn").disabled = false;
    console.log("✅ WASM Engine loaded successfully!");
}).catch(err => {
    console.error("❌ Failed to load WASM:", err);
    document.getElementById("status").innerText = "Failed to load engine.";
});

document.getElementById("solveBtn").addEventListener("click", () => {
    const scramble = document.getElementById("scramble").value;
    const t0 = performance.now();
    
    const solution = window.kociembaSolve(scramble);
    
    const t1 = performance.now();
    document.getElementById("result-box").innerHTML = 
        `<strong>Solution:</strong><br><br>${solution}<br><br><em>Calculated in ${(t1 - t0).toFixed(2)} ms</em>`;
});

console.log("🎨 Initializing 3D Scene...");
const container = document.getElementById('3d-container');

if (!container) {
    console.error("❌ ERROR: Could not find #3d-container in HTML!");
}

const scene = new THREE.Scene();

const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
renderer.setPixelRatio(window.devicePixelRatio);
container.appendChild(renderer.domElement);

const camera = new THREE.PerspectiveCamera(45, 1, 0.1, 100);
camera.position.set(6, 5, 8);
camera.lookAt(0, 0, 0);

function resizeCanvas() {
    const width = container.clientWidth;
    const height = container.clientHeight;
    renderer.setSize(width, height);
    camera.aspect = width / height;
    camera.updateProjectionMatrix();
}
resizeCanvas();
window.addEventListener('resize', resizeCanvas);

const ambientLight = new THREE.AmbientLight(0xffffff, 0.7);
scene.add(ambientLight);
const dirLight = new THREE.DirectionalLight(0xffffff, 0.8);
dirLight.position.set(10, 20, 10);
scene.add(dirLight);

const Colors = {
    R: 0xb71234, L: 0xff5800, U: 0xffffff, 
    D: 0xffd500, F: 0x009b48, B: 0x0046ad, Core: 0x111111
};

const cubies = [];
const geometry = new THREE.BoxGeometry(0.95, 0.95, 0.95);

for (let x = -1; x <= 1; x++) {
    for (let y = -1; y <= 1; y++) {
        for (let z = -1; z <= 1; z++) {
            const materials = [
                new THREE.MeshStandardMaterial({ color: x === 1 ? Colors.R : Colors.Core }),
                new THREE.MeshStandardMaterial({ color: x === -1 ? Colors.L : Colors.Core }),
                new THREE.MeshStandardMaterial({ color: y === 1 ? Colors.U : Colors.Core }),
                new THREE.MeshStandardMaterial({ color: y === -1 ? Colors.D : Colors.Core }),
                new THREE.MeshStandardMaterial({ color: z === 1 ? Colors.F : Colors.Core }),
                new THREE.MeshStandardMaterial({ color: z === -1 ? Colors.B : Colors.Core })
            ];

            const cubie = new THREE.Mesh(geometry, materials);
            cubie.position.set(x, y, z);
            scene.add(cubie);
            cubies.push(cubie);
        }
    }
}
console.log(`✅ 3D Cubies generated: ${cubies.length}`);

const cubeGroup = new THREE.Group();
cubies.forEach(c => cubeGroup.add(c));
scene.add(cubeGroup);

function animate() {
    requestAnimationFrame(animate);
    cubeGroup.rotation.y += 0.01;
    cubeGroup.rotation.x += 0.005;
    renderer.render(scene, camera);
}
animate();
console.log("✅ Render loop started!");
