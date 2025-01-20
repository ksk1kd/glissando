'use client'

import { useCallback, useEffect, useState } from 'react'

import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { API_RESOURCE_LIST_GET } from '@/constants/api'
import { AxisType } from '@/types/axis'
import { buildPublicApiUrl } from '@/utils/buildApiUrl'

type ResourcePlanningProps = {
  views: { id: string; label: string; list_name: string }[]
  members: AxisType[]
  projects: AxisType[]
  months: AxisType[]
  initialResources: number[][]
}

export function ResourcePlanningPresentation({
  views,
  members,
  projects,
  months,
  initialResources,
}: ResourcePlanningProps) {
  const [currentView, setCurrentView] = useState(views[0].id)
  const [displays, setDisplays] = useState(months)
  const [currentDisplay, setCurrentDisplay] = useState(displays[0].id)
  const [columns, setColumns] = useState(members)
  const [rows, setRows] = useState(projects)
  const [resources, setResources] = useState(initialResources)
  const [columnSums, setColumnSums] = useState(
    columns.map((_, colIndex) => resources.reduce((sum, row) => sum + row[colIndex], 0)),
  )
  const [rowSums, setRowSums] = useState(resources.map((row) => row.reduce((sum, value) => sum + value, 0)))

  const updateResources = useCallback(
    async (rows: AxisType[], columns: AxisType[]) => {
      const newResources = Array.from({ length: rows.length }, () => Array(columns.length).fill(0))
      await Promise.all(
        rows.map(async (row: AxisType, rowIndex: number) => {
          await Promise.all(
            columns.map(async (column: AxisType, columnIndex: number) => {
              let urlSearchParam = ''
              switch (currentView) {
                case 'monthly':
                  urlSearchParam = new URLSearchParams({
                    member: column.id,
                    project: row.id,
                    month: currentDisplay,
                  }).toString()
                  break
                case 'member':
                  urlSearchParam = new URLSearchParams({
                    member: currentDisplay,
                    project: row.id,
                    month: column.id,
                  }).toString()
                  break
                case 'project':
                  urlSearchParam = new URLSearchParams({
                    member: row.id,
                    project: currentDisplay,
                    month: column.id,
                  }).toString()
                  break
              }
              const resources = await (
                await fetch(
                  buildPublicApiUrl({
                    path: API_RESOURCE_LIST_GET ?? '',
                  }) +
                    '?' +
                    urlSearchParam,
                )
              ).json()
              if (resources.items.length) {
                newResources[rowIndex][columnIndex] = resources.items[0].time
              }
            }),
          )
        }),
      )
      setResources(newResources)
      setColumnSums(columns.map((_, colIndex) => newResources.reduce((sum, row) => sum + row[colIndex], 0)))
      setRowSums(newResources.map((row) => row.reduce((sum, value) => sum + value, 0)))
    },
    [currentView, currentDisplay],
  )

  useEffect(() => {
    ;(async () => {
      const selectedView = views.find((view) => view.id === currentView)
      if (selectedView) {
        setDisplays(eval(selectedView.list_name))
        setCurrentDisplay(eval(selectedView.list_name)[0].id)
        let newRows: AxisType[] = []
        let newColumns: AxisType[] = []
        switch (currentView) {
          case 'monthly':
            newRows = projects
            newColumns = members
            break
          case 'member':
            newRows = projects
            newColumns = months
            break
          case 'project':
            newRows = members
            newColumns = months
            break
        }
        setRows(newRows)
        setColumns(newColumns)
      }
    })()
  }, [currentView, views, members, projects, months])

  useEffect(() => {
    ;(async () => {
      updateResources(rows, columns)
    })()
  }, [currentDisplay, rows, columns, updateResources])

  const handleInputChange = (rowIndex: number, colIndex: number, value: number) => {
    const newResources = [...resources]
    newResources[rowIndex][colIndex] = value || 0
    setResources(newResources)
    setColumnSums(columns.map((_, colIndex) => newResources.reduce((sum, row) => sum + row[colIndex], 0)))
    setRowSums(newResources.map((row) => row.reduce((sum, value) => sum + value, 0)))
  }

  return (
    <>
      <main className="w-[1000px] mx-auto p-10">
        <header className="flex justify-center items-center gap-4 mb-10">
          <Select value={currentView} onValueChange={setCurrentView}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select View" />
            </SelectTrigger>
            <SelectContent>
              {views.map((view, index) => (
                <SelectItem value={view.id} key={view.id}>
                  {view.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <Select value={currentDisplay} onValueChange={setCurrentDisplay}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select Display" />
            </SelectTrigger>
            <SelectContent>
              {displays.map((display) => (
                <SelectItem value={display.id} key={display.id}>
                  {display.name}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </header>
        <table className="table-auto w-full">
          <thead>
            <tr>
              <th></th>
              {columns.map((column) => (
                <th key={column.id} className="p-1 w-20">
                  {column.name}
                </th>
              ))}
              <th>Total</th>
            </tr>
          </thead>
          <tbody>
            {rows.map((row, rowIndex) => (
              <tr key={row.id}>
                <th className="p-1 w-40 text-right">{row.name}</th>
                {resources[rowIndex]?.map((value, colIndex) => (
                  <td key={`${rowIndex}-${colIndex}`} className="p-1">
                    <Input
                      className={`text-center ${value !== 0 ? 'bg-fuchsia-950' : ''}`}
                      type="number"
                      value={value}
                      onChange={(e) => handleInputChange(rowIndex, colIndex, parseInt(e.target.value))}
                    />
                  </td>
                ))}
                <td className="py-1 px-6 text-center bg-violet-950">{rowSums[rowIndex]}</td>
              </tr>
            ))}
            <tr>
              <th className="p-1 w-40 text-right">Total</th>
              {columnSums.map((sum, index) => (
                <td key={index} className="py-2 px-3 text-center bg-violet-950">
                  {sum}
                </td>
              ))}
              <td className="p-1 text-right"></td>
            </tr>
          </tbody>
        </table>
      </main>
    </>
  )
}
