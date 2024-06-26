import instance from '@/utils/axios'

export interface WebsiteConfig {
  website_name: string
  website_icon: string
  website_owner: string
  website_owner_profile: string
  website_owner_avatar: string
  website_runtime: number
  website_records: string[]
}

export interface SeoConfig {
  title: string
  description: string
  og_title: string
  og_image: string
  keywords: string
  author: string
  robots: string
}

export interface CommentConfig {
  enable_comment: boolean
}

export interface EmailConfig {
  host: string
  port: number
  username: string
  password: string
  email: string
}

export interface NoticeConfig {
  title: string
  content: string
  enabled: boolean
  publish_time: number
}

export interface FrontPostCountConfig {
  count: number
}

export interface PayConfig {
  name: string
  image: string
}

export interface PayConfigRequest {
  name: string
  image: string
}

export interface SocialConfig {
  id: string
  social_name: string
  social_value: string
  css_class: string
  is_link: boolean
}

export interface SocialConfigReq {
  social_name: string
  social_value: string
  css_class: string
  is_link: boolean
}

export interface WebsiteConfigRequest {
  website_name: string
  website_icon: string
  website_owner: string
  website_owner_profile: string
  website_owner_avatar: string
  website_runtime: number
}

export interface SeoConfigRequest {
  title?: string
  description: string
  og_title?: string
  og_image?: string
  keywords: string
  author: string
  robots: string
}

export interface WebsiteConfigVO {
  website_name: string
  website_icon: string
  website_owner: string
  website_owner_profile: string
  website_owner_avatar: string
  WebsiteRuntime: number
  website_records?: string[]
}

export interface WebsiteConfigMetaVO {
  website_name: string
  website_icon: string
}

export const GetWebSiteMeta = () => {
  return instance({
    url: '/configs/website/meta',
    method: 'get'
  })
}

export const GetWebSite = () => {
  return instance({
    url: '/configs/website',
    method: 'get'
  })
}

export const UpdateWebSite = (req: WebsiteConfigRequest) => {
  return instance({
    url: '/configs/website',
    method: 'put',
    data: req
  })
}

export const AddRecord = (record: string) => {
  return instance({
    url: '/configs/website/records',
    method: 'post',
    data: {
      website_record: record
    }
  })
}

export const DeleteRecord = (record: string) => {
  return instance({
    url: '/configs/website/records?website_record=' + record,
    method: 'delete'
  })
}

export const GetSeo = () => {
  return instance({
    url: '/configs/seo',
    method: 'get'
  })
}

export const UpdateSeo = (req: SeoConfigRequest) => {
  return instance({
    url: '/configs/seo',
    method: 'put',
    data: req
  })
}

export const GetComment = () => {
  return instance({
    url: '/configs/comment',
    method: 'get'
  })
}

export interface CommentConfigRequest {
  enable_comment: boolean
}

export const UpdateComment = (req: CommentConfigRequest) => {
  return instance({
    url: '/configs/comment',
    method: 'put',
    data: req
  })
}

export const GetFriend = () => {
  return instance({
    url: '/configs/friend',
    method: 'get'
  })
}

export interface FriendConfigSwitchRequest {
  enable_friend_commit: boolean
}

export interface FriendConfigVO {
  enable_friend_commit: boolean
  introduction: string
}

export const UpdateFriendSwitch = (req: FriendConfigSwitchRequest) => {
  return instance({
    url: '/configs/friend/switch',
    method: 'put',
    data: req
  })
}

export interface FriendConfigIntroductionRequest {
  introduction: string
}

export const UpdateFriendIntroduction = (req: FriendConfigIntroductionRequest) => {
  return instance({
    url: '/configs/friend/introduction',
    method: 'put',
    data: req
  })
}

export const GetEmail = () => {
  return instance({
    url: '/configs/email',
    method: 'get'
  })
}

export interface EmailConfigRequest {
  host: string
  port: number
  username: string
  password: string
  email: string
}

export const UpdateEmail = (req: EmailConfigRequest) => {
  return instance({
    url: '/configs/email',
    method: 'put',
    data: req
  })
}

export const GetNotice = () => {
  return instance({
    url: '/configs/notice',
    method: 'get'
  })
}

export interface NoticeConfigRequest {
  title: string
  content: string
}

export const UpdateNotice = (req: NoticeConfigRequest) => {
  return instance({
    url: '/configs/notice',
    method: 'put',
    data: req
  })
}

export const UpdateNoticeEnabled = (enabled: boolean) => {
  return instance({
    url: '/configs/notice/enabled',
    method: 'put',
    data: {
      enabled: enabled
    }
  })
}

export const GetFrontPostCount = () => {
  return instance({
    url: '/configs/front-post-count',
    method: 'get'
  })
}

export interface FrontPostCountConfigRequest {
  count: number
}

export const UpdateFrontPostCount = (req: FrontPostCountConfigRequest) => {
  return instance({
    url: '/configs/front-post-count',
    method: 'put',
    data: req
  })
}

export const GetPay = () => {
  return instance({
    url: '/configs/pay',
    method: 'get'
  })
}

export const AddPay = (req: PayConfigRequest) => {
  return instance({
    url: '/configs/pay',
    method: 'post',
    data: req
  })
}

export const DeletePay = (name: string, image: string) => {
  return instance({
    url: `/configs/pay/${name}?image=${image}`,
    method: 'delete'
  })
}

export const GetSocial = () => {
  return instance({
    url: '/configs/social',
    method: 'get'
  })
}

export interface SocialConfigRequest {
  social_name: string
  social_value: string
  css_class: string
  is_link: boolean
}

export const AddSocial = (req: SocialConfigRequest) => {
  return instance({
    url: '/configs/social',
    method: 'post',
    data: req
  })
}

export const DeleteSocial = (id: string) => {
  return instance({
    url: `/configs/social/${id}`,
    method: 'delete'
  })
}

export interface SocialConfigRequest {}

export const UpdateSocial = (id: string, req: SocialConfigRequest) => {
  return instance({
    url: `/configs/social/${id}`,
    method: 'put',
    data: req
  })
}

export interface InitReq {
  website_name: string
  website_icon: string
  website_owner: string
  website_owner_profile: string
  website_owner_avatar: string
  email_server: {
    host: string
    port: number
    username: string
    password: string
    email: string
  }
  admin: {
    username: string
    password: string
  }
}

export const isInit = () => {
  return instance({
    url: '/configs/check-initialization',
    method: 'get'
  })
}

export const Init = (req: InitReq) => {
  return instance({
    url: '/configs/initialization',
    method: 'post',
    data: req
  })
}

export interface ThirdPartySiteVerification {
  key: string
  value: string
  description: string
}

export interface ThirdPartySiteVerificationRequest {
  key: string
  value: string
  description: string
}

export const GetThirdPartySiteVerification = () => {
  return instance({
    url: '/configs/third-party-site-verification',
    method: 'get'
  })
}

export const AddThirdPartySiteVerification = (req: ThirdPartySiteVerificationRequest) => {
  return instance({
    url: '/configs/third-party-site-verification',
    method: 'post',
    data: req
  })
}

export const DeleteThirdPartySiteVerification = (id: string) => {
  return instance({
    url: '/configs/third-party-site-verification/' + id,
    method: 'delete'
  })
}

export interface BaiduPushConfig {
  site: string
  token: string
}

export const GetBaiduPushConfig = () => {
  return instance({
    url: '/configs/post-index/baidu',
    method: 'get'
  })
}

export const UpdateBaiduPushConfig = (site: string, token: string) => {
  return instance({
    url: '/configs/post-index/baidu',
    method: 'put',
    data: {
      site: site,
      token: token
    }
  })
}

export interface CarouselVO {
  id: string
  title: string
  summary: string
  cover_img: string
  color: string
  created_at: number
  updated_at: number
  show: boolean
}

export interface CarouselRequest {
  id: string
  title: string
  summary: string
  cover_img: string
  color: string
  show: boolean
}

export const GetCarousel = () => {
  return instance({
    url: '/configs/carousel',
    method: 'get'
  })
}

export const AddCarousel = (data: CarouselRequest) => {
  return instance({
    url: '/configs/carousel',
    method: 'post',
    data: data
  })
}

export const UpdateCarousel = (id: string, data: CarouselRequest) => {
  return instance({
    url: `/configs/carousel/${id}`,
    method: 'put',
    data: data
  })
}

export const DeleteCarousel = (id: string) => {
  return instance({
    url: `/configs/carousel/${id}`,
    method: 'delete'
  })
}

export const ChangeCarouselShowStatus = (id: string, show: boolean) => {
  return instance({
    url: `/configs/carousel/${id}/show`,
    method: 'put',
    data: {
      show: show
    }
  })
}
