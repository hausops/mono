import {propertySvc} from '@/services/property';

type Params = {id: string};

export async function DELETE(_: Request, {params}: {params: Params}) {
  const deleted = await propertySvc.delete(params.id);
  return Response.json(deleted);
}

export async function GET(_: Request, {params}: {params: Params}) {
  const property = await propertySvc.getById(params.id);
  return Response.json(property);
}

export async function PATCH(req: Request, {params}: {params: Params}) {
  const d = await req.json();
  const updated = await propertySvc.update(params.id, d);
  return Response.json(updated);
}
