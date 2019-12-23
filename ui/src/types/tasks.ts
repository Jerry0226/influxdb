import {Task as TaskAPI} from '@influxdata/influx'
import {Task as TaskBase} from 'src/client'
import {Labels} from 'src/types'

export const TaskStatus = TaskAPI.StatusEnum

export interface Task extends TaskBase {
  labels: Labels
}
