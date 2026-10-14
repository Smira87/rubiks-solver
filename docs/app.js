console.log("🚀 app.js is starting...");

// --- 1. WASM INITIALIZATION ---
const go = new Go();
WebAssembly.instantiateStreaming(fetch("solver.wasm"), go.importObject).then((result) => {
    go.run(result.instance); 
    
    document.getElementById("status").innerText = "Morse Code ..- ._ 📟";
    document.getElementById("status").style.color = "#ff9800";
    document.getElementById("solveBtn").disabled = false;
    console.log("✅ WASM Engine loaded successfully!");
}).catch(err => {
    console.error("❌ Failed to load WASM:", err);
    document.getElementById("status").innerText = "Failed to load engine.";
});


// --- 2. THREE.JS SCENE SETUP ---
const container = document.getElementById('3d-container');
const scene = new THREE.Scene();

const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
renderer.setPixelRatio(window.devicePixelRatio);
container.appendChild(renderer.domElement);

const camera = new THREE.PerspectiveCamera(45, 1, 0.1, 100);
camera.position.set(4.2, 3.5, 5.5);
camera.lookAt(0, 0, 0);

// OrbitControls (Mouse drag)
const controls = new THREE.OrbitControls(camera, renderer.domElement);
controls.enablePan = false;
controls.enableDamping = true;
controls.dampingFactor = 0.05;
controls.enableZoom = false;

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


// --- 3. ANIMATION SYSTEM ---
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
    if (moveQueue.length === 0) {
        // If queue is empty, enable buttons
        document.getElementById("solveBtn").disabled = false;
        document.getElementById("randomBtn").disabled = false;
        return;
    }
    if (isAnimating) return;

    isAnimating = true;
    
    // Disable buttons while animating
    document.getElementById("solveBtn").disabled = true;
    document.getElementById("randomBtn").disabled = true;

    const moveStr = moveQueue.shift();
    const baseMove = moveStr[0]; 
    let modifier = moveStr[1] || ''; 
    
    const def = faceDefinitions[baseMove];
    let targetAngle = def.angle;
    
    if (modifier === "'") targetAngle *= -1;
    if (modifier === "2") targetAngle *= 2;

    const activeCubies = [];
    cubies.forEach(cubie => {
        const pos = Math.round(cubie.getWorldPosition(new THREE.Vector3())[def.axis]);
        if (pos === def.sign) activeCubies.push(cubie);
    });

    const pivot = new THREE.Object3D();
    pivot.rotation.set(0, 0, 0);
    pivot.updateMatrixWorld();
    scene.add(pivot);

    activeCubies.forEach(cubie => pivot.attach(cubie));

    let currentAngle = 0;
    const step = (targetAngle > 0 ? animationSpeed : -animationSpeed);

    function animateRotation() {
        if (Math.abs(currentAngle) < Math.abs(targetAngle)) {
            currentAngle += step;
            if (Math.abs(currentAngle) > Math.abs(targetAngle)) currentAngle = targetAngle;
            pivot.rotation[def.axis] = currentAngle;
            requestAnimationFrame(animateRotation);
        } else {
            pivot.rotation[def.axis] = targetAngle;
            pivot.updateMatrixWorld();
            
            activeCubies.forEach(cubie => {
                cubeGroup.attach(cubie);
                cubie.position.x = Math.round(cubie.position.x);
                cubie.position.y = Math.round(cubie.position.y);
                cubie.position.z = Math.round(cubie.position.z);
                cubie.rotation.x = Math.round(cubie.rotation.x / (Math.PI/2)) * (Math.PI/2);
                cubie.rotation.y = Math.round(cubie.rotation.y / (Math.PI/2)) * (Math.PI/2);
                cubie.rotation.z = Math.round(cubie.rotation.z / (Math.PI/2)) * (Math.PI/2);
            });

            scene.remove(pivot);
            isAnimating = false;
            
            setTimeout(processQueue, 30); // Pause between moves
        }
    }
    animateRotation();
}

// --- 4. APP LOGIC & UI BINDINGS ---

const scrambleInput = document.getElementById("scramble-input");
let isSolved = true; // Tracks if the physical 3D cube is in a solved state

// Utility: append move to text input
function appendToScrambleStr(move) {
    if (isSolved) {
        scrambleInput.value = move;
        isSolved = false;
    } else {
        const current = scrambleInput.value.trim();
        scrambleInput.value = current ? current + " " + move : move;
    }
}

// Random Scramble Generator
const allMoves = ["U", "U'", "U2", "D", "D'", "D2", "R", "R'", "R2", "L", "L'", "L2", "F", "F'", "F2", "B", "B'", "B2"];

document.getElementById("randomBtn").addEventListener("click", () => {
    if (isAnimating) return;

    const length = parseInt(document.getElementById("scramble-length").value) || 20;
    let sequence = [];
    let lastFace = "";

    for (let i = 0; i < length; i++) {
        let randomMove;
        let face;
        do {
            randomMove = allMoves[Math.floor(Math.random() * allMoves.length)];
            face = randomMove[0];
        } while (face === lastFace); // Prevent redundant moves like U U'
        
        lastFace = face;
        sequence.push(randomMove);
        appendToScrambleStr(randomMove);
    }

    moveQueue.push(...sequence);
    processQueue();
});

// Manual Move Buttons
document.querySelectorAll(".manual-move-btn").forEach(btn => {
    btn.addEventListener("click", (e) => {
        const move = e.target.getAttribute("data-move");
        appendToScrambleStr(move);
        moveQueue.push(move);
        processQueue();
    });
});

// Solve Button
document.getElementById("solveBtn").addEventListener("click", () => {
    if (isAnimating || moveQueue.length > 0) return;

    const scramble = scrambleInput.value.trim();
    if (!scramble) {
        document.getElementById("result-box").innerHTML = `<span style="color:#aaa">Cube is already solved!</span>`;
        return;
    }

    const t0 = performance.now();
    const solutionStr = window.kociembaSolve(scramble);
    const t1 = performance.now();

    if (solutionStr.startsWith("Error")) {
        document.getElementById("result-box").innerHTML = `<span style="color:red">${solutionStr}</span>`;
        return;
    }

    // Process solution
    const solutionMoves = solutionStr.split(' ');
    
    // Animate solution
    moveQueue.push(...solutionMoves);
    processQueue();

    document.getElementById("result-box").innerHTML = 
        `<strong>Solution (${solutionMoves.length} moves):</strong><br><br>${solutionStr}<br><br><em>Calculated in ${(t1 - t0).toFixed(2)} ms</em>`;
    
    // Clear the input because the physical cube will return to solved state
    scrambleInput.value = "";
    isSolved = true; 
});

// Render loop
function animate() {
    requestAnimationFrame(animate);
    controls.update(); 
    renderer.render(scene, camera);
}
animate();
