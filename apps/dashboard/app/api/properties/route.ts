import {propertySvc} from '@/services/property';

export async function POST(req: Request) {
  const d = await req.json();
  const created = await propertySvc.add(d);
  return Response.json(created);
}
