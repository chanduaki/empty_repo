// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'

import {
  isErrorEvent,
  isDeviceJoinEvent,
  isCRUDEvent,
  isDeviceUplinkEvent,
  isDeviceDownlinkEvent,
} from '../shared/utils/types'

import ErrorEvent from './event-types/error-event'
import CRUDEvent from './event-types/crud-event'
import DefaultEvent from './event-types/default-event'
import DeviceUplinkEvent from './event-types/device-uplink-event'
import DeviceJoinEvent from './event-types/device-join-event'
import DeviceDownlinkEvent from './event-types/device-downlink-event'

const renderApplicationEvent = (event, widget = false) => {
  if (isErrorEvent(event)) {
    return <ErrorEvent event={event} widget={widget} />
  }

  if (isCRUDEvent(event)) {
    return <CRUDEvent event={event} widget={widget} />
  }

  if (isDeviceJoinEvent(event)) {
    return <DeviceJoinEvent event={event} widget={widget} />
  }

  if (isDeviceUplinkEvent(event)) {
    return <DeviceUplinkEvent event={event} widget={widget} />
  }

  if (isDeviceDownlinkEvent(event)) {
    return <DeviceDownlinkEvent event={event} widget={widget} />
  }

  return <DefaultEvent event={event} widget={widget} />
}

export default renderApplicationEvent
