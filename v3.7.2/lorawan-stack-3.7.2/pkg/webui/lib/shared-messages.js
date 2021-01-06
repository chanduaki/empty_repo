// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

import { defineMessages } from 'react-intl'

export default defineMessages({
  // Keep these sorted alphabetically
  add: 'Add',
  addApiKey: 'Add API Key',
  addApplication: 'Add Application',
  addAttributes: 'Add Attributes',
  addCollaborator: 'Add Collaborator',
  addDevice: 'Add Device',
  addDeviceBulk: 'Device Bulk Creation',
  addGateway: 'Add Gateway',
  addOrganization: 'Add Organization',
  addPubsub: 'Add PubSub',
  address: 'Address',
  addressPlaceholder: 'host',
  addWebhook: 'Add Webhook',
  admin: 'Admin',
  all: 'All',
  altitude: 'Altitude',
  altitudeDesc: 'The altitude in meters, where 0 means sea level',
  antennas: 'Antennas',
  apiKey: 'API Key',
  apiKeyCounted: '{count, plural, one {API Key} other {API Keys}}',
  apiKeyNamePlaceholder: 'My new API Key',
  apiKeys: 'API Keys',
  appEUI: 'AppEUI',
  appId: 'Application ID',
  appKey: 'AppKey',
  application: 'Application',
  applications: 'Applications',
  applicationServerAddress: 'Application Server Address',
  approve: 'Approve',
  appSKey: 'AppSKey',
  attributeDescription:
    'Attributes can be used to set arbitrary information about the application, to be used by scripts, or simply for your own organization.',
  attributes: 'Attributes',
  automaticUpdates: 'Automatic Updates',
  brand: 'Brand',
  cancel: 'Cancel',
  changeLocation: 'Change location settings',
  changePassword: 'Change Password',
  channel: 'Channel',
  clear: 'Clear',
  claiming: 'Claiming',
  claimAuthCode: 'Claim Authentication Code',
  collaborator: 'Collaborator',
  collaboratorCounted: '{count, plural, one {Collaborator} other {Collaborators}}',
  collaboratorDeleteSuccess: 'Successfully removed collaborator',
  collaboratorEdit: 'Edit {collaboratorId}',
  collaboratorEditRights: 'Edit rights of {collaboratorId}',
  collaboratorId: 'Collaborator ID',
  collaboratorIdPlaceholder: 'collaborator-id',
  collaboratorModalWarning: 'Are you sure you want to remove {collaboratorId} as a collaborator?',
  collaboratorRemove: 'Collaborator Remove',
  collaborators: 'Collaborators',
  collaboratorUpdateSuccess: 'Successfully updated collaborator rights',
  componentApplicationServer: 'Application Server',
  componentGatewayServer: 'Gateway Server',
  componentIdentityServer: 'Identity Server',
  componentJoinServer: 'Join Server',
  componentNetworkServer: 'Network Server',
  confirmPassword: 'Confirm Password',
  connected: 'Connected',
  connecting: 'Connecting',
  createApiKey: 'Create API Key',
  created: 'Created',
  createdAt: 'Created at',
  currentCollaborators: 'Current Collaborators',
  data: 'Data',
  defineRights: 'Define Rights',
  description: 'Description',
  devAddr: 'Device Address',
  devDesc: 'Device Description',
  devEUI: 'DevEUI',
  deviceCounted: '{count, plural, one {Device} other {Devices}}',
  devices: 'Devices',
  devID: 'Device ID',
  devName: 'Device Name',
  disabled: 'Disabled',
  disconnected: 'Disconnected',
  downlink: 'Downlink',
  downlinkAck: 'Downlink Ack',
  downlinkFailed: 'Downlink Failed',
  downlinkNack: 'Downlink Nack',
  downlinkPush: 'Downlink Push',
  downlinkQueued: 'Downlink Queued',
  downlinkReplace: 'Downlink Replace',
  downlinkSent: 'Downlink Sent',
  downlinksScheduled: 'Downlinks (re)scheduled',
  edit: 'Edit',
  email: 'Email',
  emailAddress: 'Email Address',
  empty: 'Empty',
  enabled: 'Enabled',
  entityId: 'Entity ID',
  eventsCannotShow: 'Cannot show events',
  fetching: 'Fetching data…',
  firmwareVersion: 'Firmware Version',
  frequencyPlan: 'Frequency Plan',
  fNwkSIntKey: 'FNwkSIntKey',
  nwkSKey: 'NwkSKey',
  gatewayDescription: 'Gateway Description',
  gatewayEUI: 'Gateway EUI',
  gatewayID: 'Gateway ID',
  gatewayLocation: 'Gateway Location',
  gatewayName: 'Gateway Name',
  gateways: 'Gateways',
  gatewayScheduleDownlinkLate: 'Schedule Downlink Late',
  gatewayServerAddress: 'Gateway Server Address',
  gatewayStatus: 'Gateway Status',
  gatewayUpdateOptions: 'Gateway Updates',
  general: 'General',
  generalInformation: 'General Information',
  generalSettings: 'General Settings',
  hardware: 'Hardware',
  hardwareVersion: 'Hardware version',
  id: 'ID',
  import: 'Import',
  importDevices: 'Import Devices',
  integrations: 'Integrations',
  joinAccept: 'Join Accept',
  joinEUI: 'JoinEUI',
  joinServerAddress: 'Join Server Address',
  key: 'key',
  keyEdit: 'Edit API Key',
  keyId: 'Key ID',
  lastSeen: 'Last seen',
  latitude: 'Latitude',
  latitudeDesc: 'The North-South position in degrees, where 0 is the equator',
  link: 'Link',
  linked: 'Linked',
  location: 'Location',
  locationSolved: 'Location Solved',
  login: 'Login',
  logout: 'Logout',
  longitude: 'Longitude',
  longitudeDesc: 'The East-West position in degrees, where 0 is the Prime Meridian (Greenwich)',
  lorawanInformation: 'LoRaWAN Information',
  lorawanOptions: 'LoRaWAN Options',
  macVersion: 'MAC Version',
  messageTypes: 'Message types',
  model: 'Model',
  mqtt: 'MQTT',
  name: 'Name',
  networkServerAddress: 'Network Server Address',
  noDesc: 'This device has no description',
  noEvents: '{entityId} has not sent any events recently',
  noLocation: 'No location information available',
  noMatch: 'No items found',
  none: 'None',
  notAvailable: 'n/a',
  notLinked: 'Not linked',
  notSet: 'Not set',
  nsAddress: 'Network Server Address',
  nsEmptyDefault: 'Leave empty to link to the Network Server in the same cluster',
  nwkSEncKey: 'NwkSEncKey',
  nwkKey: 'NwkKey',
  ok: 'Ok',
  offline: 'Offline',
  online: 'Online',
  options: 'Options',
  organization: 'Organization',
  organizationId: 'Organization ID',
  organizations: 'Organizations',
  otherCluster: 'Other Cluster',
  overview: 'Overview',
  password: 'Password',
  pause: 'Pause',
  payloadFormatters: 'Payload Formatters',
  payloadFormattersDownlink: 'Downlink Payload Formatters',
  payloadFormattersUpdateFailure: 'There was an error updating the Payload Formatter',
  payloadFormattersUpdateSuccess: 'Payload Formatter has been set successfully',
  payloadFormattersUplink: 'Uplink Payload Formatters',
  phyVersion: 'PHY Version',
  port: 'Port',
  privacyPolicy: 'Privacy Policy',
  provider: 'Provider',
  provisionedOnExternalJoinServer: 'Provisioned on external Join Server',
  public: 'Public',
  pubsubBaseTopic: 'Base Topic',
  pubsubFormat: 'PubSub Format',
  pubsubId: 'PubSub ID',
  pubsubs: 'PubSubs',
  refresh: 'Refresh',
  refreshPage: 'Refresh page',
  removeCollaborator: 'Remove Collaborator',
  restartStream: 'Restart Stream',
  resume: 'Resume',
  rights: 'Rights',
  saveChanges: 'Save Changes',
  searchById: 'Search by ID',
  secure: 'Secure',
  settings: 'Settings',
  sNwkSIKey: 'SNwkSIntKey',
  stable: 'Stable',
  state: 'State',
  stateApproved: 'Approved',
  stateFlagged: 'Flagged',
  stateRejected: 'Rejected',
  stateRequested: 'Requested',
  stateSuspended: 'Suspended',
  status: 'Status',
  statusUnknown: 'Status Unknown',
  takeMeBack: 'Take me back',
  termsAndCondition: 'Terms and Conditions',
  time: 'Time',
  traffic: 'Traffic',
  type: 'Type',
  unknown: 'Unknown',
  updatedAt: 'Last updated at',
  uplink: 'Uplink',
  uplinkMessage: 'Uplink Message',
  uplinksReceived: 'Uplinks Received',
  user: 'User',
  userDelete: 'Delete User',
  userEdit: 'Edit User',
  userId: 'User ID',
  userManagement: 'User Management',
  username: 'Username',
  users: 'Users',
  validateAddress: 'Wrong address format',
  validateAddressFormat: 'Wrong address format, should be host or host:port',
  validateDelayFormat: 'Delay must be a valid, positive number',
  validateEmail: 'Invalid email address',
  validateFormat: 'Wrong format',
  validateIdFormat: 'This value may only contain lowercase letters, numbers and dashes (-)',
  validateInt32: 'Invalid numeric format',
  validateLatLong: 'Invalid value format (valid format e.g. 52.15249000)',
  validateMqttPassword: 'Password should be empty or at least 2 characters',
  validateMqttUrl: 'Invalid MQTT URL format',
  validateNoSpaces: 'Invalid format, should contain no spaces',
  validatePasswordMatch: 'Passwords should match',
  validateRequired: 'This field is required',
  validateRights: 'At least one right should be selected',
  validateTooLong: 'Too long',
  validateTooShort: 'Too short',
  validateUrl: 'Invalid URL format',
  validFrom: 'Valid From',
  validTo: 'Valid To',
  value: 'value',
  webhookBaseUrl: 'Base URL',
  webhookFormat: 'Webhook Format',
  webhookId: 'Webhook ID',
  webhooks: 'Webhooks',
})