const { Kafka } = require("kafkajs");

async function run() {
  try {
    const kafka = new Kafka({
      clientId: "myapp",
      brokers: ["localhost:9092"],
    });

    const admin = kafka.admin();
    console.log("Connecting..");
    await admin.connect();
    console.log("Connected!");
    // Topic paritions from A-M, N-Z
    await admin.createTopics({
      topics: [
        {
          topic: "Users",
          numPartitions: 2,
        },
      ],
    });
    console.log("Topics created successfully");
    await admin.disconnect();
  } catch (e) {
    console.error(`Something bad happened ${e}`);
  } finally {
    process.exit(0);
  }
}

run();
