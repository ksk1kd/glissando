import { API_MEMBER_LIST_GET, API_PROJECT_LIST_GET } from '@/constants/api'
import { buildApiUrl } from '@/utils/buildApiUrl'

import { ResourcePlanningPresentation } from './presentational'

export async function ResourcePlanningContainer() {
  const members = await (await fetch(buildApiUrl({ path: API_MEMBER_LIST_GET ?? '' }))).json()
  const projects = await (await fetch(buildApiUrl({ path: API_PROJECT_LIST_GET ?? '' }))).json()
  const months = Array.from({ length: 12 }, (_, i) => ({
    id: new Date(0, i).toLocaleString('default', { month: 'short' }),
    name: new Date(0, i).toLocaleString('default', { month: 'short' }),
  }))

  return (
    <>
      <ResourcePlanningPresentation members={members.items} projects={projects.items} months={months} />
    </>
  )
}
