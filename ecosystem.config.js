module.exports = {
    apps: [
        {
            name: "API", // application name
            script: "go run main.go", // script path to pm2 start
            instances: 1, // number process of application
            max_memory_restart: "1G", // restart if it exceeds the amount of memory specified
        },
    ],
};