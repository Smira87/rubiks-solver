console.log("🚀 app.js is starting...");

const go = new Go();
WebAssembly.instantiateStreaming(fetch("solver.wasm"), go.importObject).then((result) => {
    go.run(result.instance); 
    
    document.getElementById("status").innerText = "Engine Ready 🚀";
    document.getElementById("status").style.color = "#4CAF50";
    document.getElementById("solveBtn").disabled = false;
}).catch(err => {
    console.error("❌ Failed to load WASM:", err);
    document.getElementById("status").innerText = "Failed to load engine.";
});

const container = document.getElementById('3d-container');
const scene = new THREE.Scene();

const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
renderer.setPixelRatio(window.devicePixelRatio);
container.appendChild(renderer.domElement);

const camera = new THREE.PerspectiveCamera(45, 1, 0.1, 100);
camera.position.set(6, 5, 8);
camera.lookAt(0, 0, 0);

// Add Mouse Controls
const controls = new THREE.OrbitControls(camera, renderer.domElement);
controls.enablePan = false;
controls.enableDamping = true;
controls.dampingFactor = 0.05;

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

const cubeGroup = new THREE.Group();
cubies.forEach(c => cubeGroup.add(c));
scene.add(cubeGroup);

// --- ANIMATION SYSTEM ---

let moveQueue = [];
let isAnimating = false;
const animationSpeed = 0.15; // Radians per frame

const faceDefinitions = {
    'R': { axis: 'x', sign: 1, angle: -Math.PI / 2 },
    'L': { axis: 'x', sign: -1, angle: Math.PI / 2 },
    'U': { axis: 'y', sign: 1, angle: -Math.PI / 2 },
    'D': { axis: 'y', sign: -1, angle: Math.PI / 2 },
    'F': { axis: 'z', sign: 1, angle: -Math.PI / 2 },
    'B': { axis: 'z', sign: -1, angle: Math.PI / 2 }
};

function processQueue() {
    if (moveQueue.length === 0 || isAnimating) return;

    isAnimating = true;
    const move = moveQueue.shift();
    executeMove(move);
}

function executeMove(moveStr) {
    const baseMove = moveStr[0]; // U, D, R, L, F, B
    let modifier = moveStr[1] || ''; // ' or 2 or nothing
    
    const def = faceDefinitions[baseMove];
    let targetAngle = def.angle;
    
    if (modifier === "'") targetAngle *= -1;
    if (modifier === "2") targetAngle *= 2;

    // Find the 9 cubies that belong to this face based on their current world position
    const activeCubies = [];
    cubies.forEach(cubie => {
        // We use Math.round to avoid floating point drift
        const pos = Math.round(cubie.getWorldPosition(new THREE.Vector3())[def.axis]);
        if (pos === def.sign) {
            activeCubies.push(cubie);
        }
    });

    // Create a pivot point at the center of the cube
    const pivot = new THREE.Object3D();
    pivot.rotation.set(0, 0, 0);
    pivot.updateMatrixWorld();
    scene.add(pivot);

    // Attach the 9 cubies to the pivot
    activeCubies.forEach(cubie => {
        pivot.attach(cubie);
    });

    let currentAngle = 0;
    const step = (targetAngle > 0 ? animationSpeed : -animationSpeed);

    function animateRotation() {
        if (Math.abs(currentAngle) < Math.abs(targetAngle)) {
            currentAngle += step;
            // Prevent overshooting
            if (Math.abs(currentAngle) > Math.abs(targetAngle)) {
                currentAngle = targetAngle;
            }
            pivot.rotation[def.axis] = currentAngle;
            requestAnimationFrame(animateRotation);
        } else {
            // Animation finished
            pivot.rotation[def.axis] = targetAngle;
            pivot.updateMatrixWorld();
            
            // Detach cubies and add them back to the main group
            activeCubies.forEach(cubie => {
                cubeGroup.attach(cubie);
                
                // Fix floating point drift
                cubie.position.x = Math.round(cubie.position.x);
                cubie.position.y = Math.round(cubie.position.y);
                cubie.position.z = Math.round(cubie.position.z);
                
                cubie.rotation.x = Math.round(cubie.rotation.x / (Math.PI/2)) * (Math.PI/2);
                cubie.rotation.y = Math.round(cubie.rotation.y / (Math.PI/2)) * (Math.PI/2);
                cubie.rotation.z = Math.round(cubie.rotation.z / (Math.PI/2)) * (Math.PI/2);
            });

            scene.remove(pivot);
            isAnimating = false;
            
            // Proceed to next move
            setTimeout(processQueue, 50); // Small pause between moves
        }
    }
    animateRotation();
}

// --- BIND UI TO ANIMATION ---

document.getElementById("solveBtn").addEventListener("click", () => {
    if (isAnimating || moveQueue.length > 0) return; // Prevent double clicks

    const scramble = document.getElementById("scramble").value.trim();
    if (!scramble) return;

    // Apply the scramble visually first
    const scrambleMoves = scramble.split(' ');
    moveQueue.push(...scrambleMoves);
    processQueue();

    const t0 = performance.now();
    const solutionStr = window.kociembaSolve(scramble);
    const t1 = performance.now();

    if (solutionStr.startsWith("Error")) {
        document.getElementById("result-box").innerHTML = `<span style="color:red">${solutionStr}</span>`;
        return;
    }

    const solutionMoves = solutionStr.split(' ');
    
    // Add solution moves to the queue to play right after scramble finishes
    moveQueue.push(...solutionMoves);
    
    document.getElementById("result-box").innerHTML = 
        `<strong>Solution (${solutionMoves.length} moves):</strong><br><br>${solutionStr}<br><br><em>Engine calc: ${(t1 - t0).toFixed(2)} ms</em>`;
});

function animate() {
    requestAnimationFrame(animate);
    controls.update(); // Required for OrbitControls damping
    renderer.render(scene, camera);
}
animate();
