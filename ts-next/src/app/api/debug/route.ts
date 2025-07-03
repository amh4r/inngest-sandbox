export function GET() {
  console.log(process.env);
  return new Response("debug");
}
