import { useRouter } from "next/router";

export default function Car() {
  const router = useRouter();
  const { id } = router.query;

  return <h1>hello {id}</h1>;
}

const getStaticProps;
export async function getStaticProps({ params }) {
  const req = await fetch("localhost:3000/cars.json");
}
