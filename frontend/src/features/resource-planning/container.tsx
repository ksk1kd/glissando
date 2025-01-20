import { API_MEMBER_LIST_GET, API_PROJECT_LIST_GET } from '@/constants/api'
import { buildApiUrl } from '@/utils/buildApiUrl'

import { ResourcePlanningPresentation } from './presentational'

export async function ResourcePlanningContainer() {
  const views = [
    {
      id: 'monthly',
      label: 'Monthly',
      list_name: 'months',
    },
    {
      id: 'member',
      label: 'Member',
      list_name: 'members',
    },
    {
      id: 'project',
      label: 'Project',
      list_name: 'projects',
    },
  ]
  const members = await (await fetch(buildApiUrl({ path: API_MEMBER_LIST_GET ?? '' }))).json()
  const projects = await (await fetch(buildApiUrl({ path: API_PROJECT_LIST_GET ?? '' }))).json()
  const months = Array.from({ length: 12 }, (_, i) => ({
    id: new Date(0, i).toLocaleString('default', { month: 'short' }),
    name: new Date(0, i).toLocaleString('default', { month: 'short' }),
  }))
  const initialResources = Array.from({ length: projects.items.length }, () => Array(members.items.length).fill(0))

  return (
    <>
      <ResourcePlanningPresentation
        views={views}
        members={members.items}
        projects={projects.items}
        months={months}
        initialResources={initialResources}
      />
    </>
  )
}
