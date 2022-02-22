const { Kafka } = require("kafkajs");
const msg = process.argv[2];

async function run() {
  try {
    const kafka = new Kafka({
      clientId: "myapp",
      brokers: ["localhost:9092"],
    });

    const producer = kafka.producer();
    console.log("Connecting..");
    await producer.connect();
    console.log("Connected!");
    const partition = msg < "N" ? 0 : 1;
    const result = await producer.send({
      topic: "Users",
      messages: [
        {
          value: msg,
          partition: partition,
        },
      ],
    });
    console.log(`Message sent successfully ${JSON.stringify(result)}`);
    await producer.disconnect();
  } catch (e) {
    console.error(`Something bad happened ${e}`);
  } finally {
    process.exit(0);
  }
}
run();
