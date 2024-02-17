async function worker(id) {
  const form = new FormData();
  form.append('user', "worker" + id);

  return fetch('http://localhost:4000/ticket', {
    method: 'POST',
    body: form,
  }).then((res) => res.json());
}

async function main() {
  const workers = []
  for (let i = 1; i <= 200; i++) {
    workers.push(worker(i))
  }
  await Promise.all(workers)
}

main();
