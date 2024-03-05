CREATE MATERIALIZED VIEW content_videos AS
SELECT  c.id                                                                                                                          AS "id"
       ,c.title                                                                                                                       AS "title"
       ,c.tagline                                                                                                                     AS "tagline"
       ,c.description                                                                                                                 AS "description"
       ,c.overview                                                                                                                    AS "overview"
       ,c.duration                                                                                                                    AS "duration"
       ,c.runtime                                                                                                                     AS "runtime"
       ,c.is_free                                                                                                                     AS isFree
       ,ct.content_type_name                                                                                                          AS contentType
       ,cst.content_subtype_name                                                                                                      AS subContentType
       ,pt.name                                                                                                                       AS publishedType
       ,TO_CHAR(c.published_date,'DD-MM-YYYY')                                                                                        AS publishedDate
       ,l.name                                                                                                                        AS "language"
       ,jsonb_build_object( 'name',pr.parental_rating_name,'shortName',pr.parental_rating_short_name,'tagline',pr.parental_rating_tagline,'description',pr.parental_rating_description,'warningTitle',pr.parental_rating_warning_title,'isKid',pr.is_kid,'image',jsonb_build_object( 'type',it.name,'profile',pi.profile,'imageLink',pi.image_link,'layoutType',pi.layout_type ) ) AS parentalRating
       ,jsonb_build_object( 'name',dl.deeplink_name,'link',dl.deeplink_link )  AS deeplink
       ,jsonb_agg( jsonb_build_object( 'type',cit.name,'profile',ci.profile,'imageLink',ci.image_link,'layoutType',ci.layout_type ) ) AS images
       ,jsonb_build_object( 'parentContentId',sm.parent_content_id,'subParentContentId',sm.sub_parent_content_id,'episodeNumber',sm.episode_number,'index',sm.index,'popularity',sm.popularity ) AS seriesMeta
       ,jsonb_build_object( 'releasedDate',rs.released_date,'OttReleasedDate',rs.ott_release_date,'tvReleasedDate',rs.tv_release_date,'budget',rs.budget,'revenue',rs.revenue ) AS releaseStatus
       ,jsonb_build_object( 'title',am.alt_meta_title,'tagline',am.alt_meta_tagline,'description',am.alt_meta_description,'overview',am.alt_meta_overview,'language',aml.display_name ) AS altMeta
       ,jsonb_agg( jsonb_build_object( 'name',dt.display_tag_name,'displayName',dt.display_tag_display_name,'carouselType',dt.carousel_type,'image',jsonb_build_object( 'type',dtit.name,'profile',dti.profile,'imageLink',dti.image_link,'layoutType',dti.layout_type ) ) ) AS displayTags
       ,jsonb_agg( jsonb_build_object( 'name',cp.name,'key',cp.key,'value',cp.value ) )                                               AS customParameters
       ,jsonb_agg( jsonb_build_object( 'Id',pa.id,'name',pa.partner_name,'displayName',pa.partner_display_name,'description',pa.partner_description,'image',jsonb_build_object( 'type',pait.name,'profile',pai.profile,'imageLink',pai.image_link,'layoutType',pai.layout_type ) ) ) AS partners
       ,jsonb_build_object( 'displayName',co.display_name,'name',co.name,'imageLink',co.image_link,'shortName',co.short_name,'phoneCode',co.phone_code,'active',co.active,'timezone',co.timezone ) AS country
       ,jsonb_build_object( 'providerName',r.provider_name,'providerDisplayName',r.provider_display_name,'rating',r.rating,'ratingLimit',r.rating_limit,'type',r.type ) AS rating
       ,jsonb_build_object( 'displayName',st.display_name,'name',st.name )                                                            AS studio
       ,jsonb_build_object( 'displayName',ge.genre_display_name,'name',ge.genre_name,'image',jsonb_build_object( 'type',geit.name,'profile',gei.profile,'imageLink',gei.image_link,'layoutType',gei.layout_type ) ) AS genre
       ,jsonb_build_object( 'name',ph.production_house_name,'displayName',ph.production_house_display_name,'description',ph.production_house_description,'image',jsonb_build_object( 'type',phit.name,'profile',phi.profile,'imageLink',phi.image_link,'layoutType',phi.layout_type ) ) AS productionHouse
       ,jsonb_agg( jsonb_build_object( 'type',pv.preview_type,'mimetype',pvmt.preview_mime_type_name,'carouselType',pv.preview_carousel_type,'image',jsonb_build_object( 'type',pvit.name,'profile',pvi.profile,'imageLink',pvi.image_link,'layoutType',pvi.layout_type ),'videoURL',jsonb_build_object( 'type',vut.video_url_type_name,'profile',vup.video_url_profile_name,'drm_enabled',vu.drm_enabled,'licenceUrl',vu.licence_url,'contentUrl',vu.content_url,'protocol',vupr.video_url_protocol_name,'encryptionType',vue.video_url_encryption_type_name,'qualityBitrate',jsonb_build_object( 'name',qb.quality_bitrate_name,'displayName',qb.quality_bitrate_display_name,'minRange',qb.min_range,'maxRange',qb.max_range ) ) ) ) AS preview
FROM content c
LEFT JOIN content_type ct
ON c.content_type_id = ct.id
LEFT JOIN content_subtype cst
ON c.content_subtype_id = cst.id
LEFT JOIN published_type pt
ON c.published_type_id = pt.id
LEFT JOIN language l
ON c.language_id = l.id
LEFT JOIN parental_rating pr
ON c.parental_rating_id = pr.id
LEFT JOIN image pi
ON pr.image_id = pi.id
LEFT JOIN image_type it
ON pi.type_id = it.id
LEFT JOIN image ci
ON ci.content_id = c.id
LEFT JOIN image_type cit
ON ci.type_id = cit.id
LEFT JOIN release_status rs
ON c.release_status_id = rs.id
LEFT JOIN alt_meta_content amc
ON c.id = amc.content_id
LEFT JOIN alt_meta am
ON amc.alt_meta_id = am.id
LEFT JOIN production_house_content phc
ON c.id = phc.content_id
LEFT JOIN production_house ph
ON phc.production_house_id = ph.id
LEFT JOIN image phi
ON ph.image_id = phi.id
LEFT JOIN image_type phit
ON phi.type_id = phit.id
LEFT JOIN local_language aml
ON am.locale_language_id = aml.id
LEFT JOIN studio_type st
ON c.studio_type_id = st.id
LEFT JOIN display_tag_content dtc
ON c.id = dtc.content_id
LEFT JOIN display_tag dt
ON dtc.display_tag_id = dt.id
LEFT JOIN image dti
ON dt.image_id = dti.id
LEFT JOIN image_type dtit
ON dti.type_id = dtit.id
LEFT JOIN rating_content rc
ON c.id = rc.content_id
LEFT JOIN rating r
ON rc.rating_id = r.id
LEFT JOIN genre_content gc
ON c.id = gc.content_id
LEFT JOIN genre ge
ON gc.genre_id = ge.id
LEFT JOIN image gei
ON ge.image_id = gei.id
LEFT JOIN image_type geit
ON gei.type_id = geit.id
LEFT JOIN deeplink dl
ON c.id = dl.content_id
LEFT JOIN custom_parameter_content cpc
ON c.id = cpc.content_id
LEFT JOIN custom_parameter cp
ON cpc.custom_parameter_id = cp.id
LEFT JOIN partner_content pac
ON c.id = pac.content_id
LEFT JOIN partner pa
ON pac.partner_id = pa.id
LEFT JOIN image pai
ON pa.image_id = pai.id
LEFT JOIN image_type pait
ON pai.type_id = pait.id
LEFT JOIN series_meta sm
ON c.series_meta_id = sm.id
LEFT JOIN country_content cc
ON c.id = cc.content_id
LEFT JOIN country co
ON cc.country_id = co.id
LEFT JOIN preview_content pvc
ON c.id = pvc.content_id
LEFT JOIN preview pv
ON pvc.preview_id = pv.id
LEFT JOIN preview_mime_type pvmt
ON pv.mime_type_id = pvmt.id
LEFT JOIN image pvi
ON pv.image_id = pvi.id
LEFT JOIN image_type pvit
ON pvi.type_id = pvit.id
LEFT JOIN video_url vu
ON pv.video_id = vu.id
LEFT JOIN video_url_type vut
ON vu.type_id = vut.id
LEFT JOIN video_url_protocol vupr
ON vu.protocol_id = vupr.id
LEFT JOIN video_url_encryption_type vue
ON vu.encryption_type_id = vue.id
LEFT JOIN video_url_profile vup
ON vu.profile_id = vup.id
LEFT JOIN quality_bitrate qb
ON vu.id = qb.video_url_id
GROUP BY  c.id
         ,ct.content_type_name
         ,cst.content_subtype_name
         ,pt.name
         ,l.name
         ,rs.released_date
         ,rs.ott_release_date
         ,rs.tv_release_date
         ,rs.budget
         ,rs.revenue
         ,am.alt_meta_title
         ,am.alt_meta_tagline
         ,am.alt_meta_description
         ,am.alt_meta_overview
         ,aml.display_name
         ,st.name
         ,st.display_name
         ,dl.deeplink_name
         ,dl.deeplink_link
         ,sm.parent_content_id
         ,sm.sub_parent_content_id
         ,sm.episode_number
         ,sm.index
         ,pr.parental_rating_name
         ,pr.parental_rating_short_name
         ,pr.parental_rating_tagline
         ,pr.parental_rating_description
         ,pr.parental_rating_warning_title
         ,pr.is_kid
         ,it.name
         ,pi.profile
         ,pi.image_link
         ,pi.layout_type
         ,co.display_name
         ,co.name
         ,co.image_link
         ,co.short_name
         ,co.phone_code
         ,co.active
         ,co.timezone
         ,ph.production_house_name
         ,ph.production_house_display_name
         ,ph.production_house_description
         ,phit.name
         ,phi.profile
         ,phi.image_link
         ,phi.layout_type
         ,ge.genre_display_name
         ,ge.genre_name
         ,geit.name
         ,gei.profile
         ,gei.image_link
         ,gei.layout_type
         ,r.provider_name
         ,r.provider_display_name
         ,r.rating
         ,r.rating_limit
         ,r.type
         ,sm.popularity;