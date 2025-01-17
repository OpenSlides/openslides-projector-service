package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Meeting struct {
	AdminGroupID                                 *int            `json:"admin_group_id"`
	AgendaEnableNumbering                        *bool           `json:"agenda_enable_numbering"`
	AgendaItemCreation                           *string         `json:"agenda_item_creation"`
	AgendaItemIDs                                []int           `json:"agenda_item_ids"`
	AgendaNewItemsDefaultVisibility              *string         `json:"agenda_new_items_default_visibility"`
	AgendaNumberPrefix                           *string         `json:"agenda_number_prefix"`
	AgendaNumeralSystem                          *string         `json:"agenda_numeral_system"`
	AgendaShowInternalItemsOnProjector           *bool           `json:"agenda_show_internal_items_on_projector"`
	AgendaShowSubtitles                          *bool           `json:"agenda_show_subtitles"`
	AgendaShowTopicNavigationOnDetailView        *bool           `json:"agenda_show_topic_navigation_on_detail_view"`
	AllProjectionIDs                             []int           `json:"all_projection_ids"`
	AnonymousGroupID                             *int            `json:"anonymous_group_id"`
	ApplauseEnable                               *bool           `json:"applause_enable"`
	ApplauseMaxAmount                            *int            `json:"applause_max_amount"`
	ApplauseMinAmount                            *int            `json:"applause_min_amount"`
	ApplauseParticleImageUrl                     *string         `json:"applause_particle_image_url"`
	ApplauseShowLevel                            *bool           `json:"applause_show_level"`
	ApplauseTimeout                              *int            `json:"applause_timeout"`
	ApplauseType                                 *string         `json:"applause_type"`
	AssignmentCandidateIDs                       []int           `json:"assignment_candidate_ids"`
	AssignmentIDs                                []int           `json:"assignment_ids"`
	AssignmentPollAddCandidatesToListOfSpeakers  *bool           `json:"assignment_poll_add_candidates_to_list_of_speakers"`
	AssignmentPollBallotPaperNumber              *int            `json:"assignment_poll_ballot_paper_number"`
	AssignmentPollBallotPaperSelection           *string         `json:"assignment_poll_ballot_paper_selection"`
	AssignmentPollDefaultBackend                 *string         `json:"assignment_poll_default_backend"`
	AssignmentPollDefaultGroupIDs                []int           `json:"assignment_poll_default_group_ids"`
	AssignmentPollDefaultMethod                  *string         `json:"assignment_poll_default_method"`
	AssignmentPollDefaultOnehundredPercentBase   *string         `json:"assignment_poll_default_onehundred_percent_base"`
	AssignmentPollDefaultType                    *string         `json:"assignment_poll_default_type"`
	AssignmentPollEnableMaxVotesPerOption        *bool           `json:"assignment_poll_enable_max_votes_per_option"`
	AssignmentPollSortPollResultByVotes          *bool           `json:"assignment_poll_sort_poll_result_by_votes"`
	AssignmentsExportPreamble                    *string         `json:"assignments_export_preamble"`
	AssignmentsExportTitle                       *string         `json:"assignments_export_title"`
	ChatGroupIDs                                 []int           `json:"chat_group_ids"`
	ChatMessageIDs                               []int           `json:"chat_message_ids"`
	CommitteeID                                  int             `json:"committee_id"`
	ConferenceAutoConnect                        *bool           `json:"conference_auto_connect"`
	ConferenceAutoConnectNextSpeakers            *int            `json:"conference_auto_connect_next_speakers"`
	ConferenceEnableHelpdesk                     *bool           `json:"conference_enable_helpdesk"`
	ConferenceLosRestriction                     *bool           `json:"conference_los_restriction"`
	ConferenceOpenMicrophone                     *bool           `json:"conference_open_microphone"`
	ConferenceOpenVideo                          *bool           `json:"conference_open_video"`
	ConferenceShow                               *bool           `json:"conference_show"`
	ConferenceStreamPosterUrl                    *string         `json:"conference_stream_poster_url"`
	ConferenceStreamUrl                          *string         `json:"conference_stream_url"`
	CustomTranslations                           json.RawMessage `json:"custom_translations"`
	DefaultGroupID                               int             `json:"default_group_id"`
	DefaultMeetingForCommitteeID                 *int            `json:"default_meeting_for_committee_id"`
	DefaultProjectorAgendaItemListIDs            []int           `json:"default_projector_agenda_item_list_ids"`
	DefaultProjectorAmendmentIDs                 []int           `json:"default_projector_amendment_ids"`
	DefaultProjectorAssignmentIDs                []int           `json:"default_projector_assignment_ids"`
	DefaultProjectorAssignmentPollIDs            []int           `json:"default_projector_assignment_poll_ids"`
	DefaultProjectorCountdownIDs                 []int           `json:"default_projector_countdown_ids"`
	DefaultProjectorCurrentListOfSpeakersIDs     []int           `json:"default_projector_current_list_of_speakers_ids"`
	DefaultProjectorListOfSpeakersIDs            []int           `json:"default_projector_list_of_speakers_ids"`
	DefaultProjectorMediafileIDs                 []int           `json:"default_projector_mediafile_ids"`
	DefaultProjectorMessageIDs                   []int           `json:"default_projector_message_ids"`
	DefaultProjectorMotionBlockIDs               []int           `json:"default_projector_motion_block_ids"`
	DefaultProjectorMotionIDs                    []int           `json:"default_projector_motion_ids"`
	DefaultProjectorMotionPollIDs                []int           `json:"default_projector_motion_poll_ids"`
	DefaultProjectorPollIDs                      []int           `json:"default_projector_poll_ids"`
	DefaultProjectorTopicIDs                     []int           `json:"default_projector_topic_ids"`
	Description                                  *string         `json:"description"`
	EnableAnonymous                              *bool           `json:"enable_anonymous"`
	EndTime                                      *int            `json:"end_time"`
	ExportCsvEncoding                            *string         `json:"export_csv_encoding"`
	ExportCsvSeparator                           *string         `json:"export_csv_separator"`
	ExportPdfFontsize                            *int            `json:"export_pdf_fontsize"`
	ExportPdfLineHeight                          *float32        `json:"export_pdf_line_height"`
	ExportPdfPageMarginBottom                    *int            `json:"export_pdf_page_margin_bottom"`
	ExportPdfPageMarginLeft                      *int            `json:"export_pdf_page_margin_left"`
	ExportPdfPageMarginRight                     *int            `json:"export_pdf_page_margin_right"`
	ExportPdfPageMarginTop                       *int            `json:"export_pdf_page_margin_top"`
	ExportPdfPagenumberAlignment                 *string         `json:"export_pdf_pagenumber_alignment"`
	ExportPdfPagesize                            *string         `json:"export_pdf_pagesize"`
	ExternalID                                   *string         `json:"external_id"`
	FontBoldID                                   *int            `json:"font_bold_id"`
	FontBoldItalicID                             *int            `json:"font_bold_italic_id"`
	FontChyronSpeakerNameID                      *int            `json:"font_chyron_speaker_name_id"`
	FontItalicID                                 *int            `json:"font_italic_id"`
	FontMonospaceID                              *int            `json:"font_monospace_id"`
	FontProjectorH1ID                            *int            `json:"font_projector_h1_id"`
	FontProjectorH2ID                            *int            `json:"font_projector_h2_id"`
	FontRegularID                                *int            `json:"font_regular_id"`
	ForwardedMotionIDs                           []int           `json:"forwarded_motion_ids"`
	GroupIDs                                     []int           `json:"group_ids"`
	ID                                           int             `json:"id"`
	ImportedAt                                   *int            `json:"imported_at"`
	IsActiveInOrganizationID                     *int            `json:"is_active_in_organization_id"`
	IsArchivedInOrganizationID                   *int            `json:"is_archived_in_organization_id"`
	JitsiDomain                                  *string         `json:"jitsi_domain"`
	JitsiRoomName                                *string         `json:"jitsi_room_name"`
	JitsiRoomPassword                            *string         `json:"jitsi_room_password"`
	Language                                     string          `json:"language"`
	ListOfSpeakersAllowMultipleSpeakers          *bool           `json:"list_of_speakers_allow_multiple_speakers"`
	ListOfSpeakersAmountLastOnProjector          *int            `json:"list_of_speakers_amount_last_on_projector"`
	ListOfSpeakersAmountNextOnProjector          *int            `json:"list_of_speakers_amount_next_on_projector"`
	ListOfSpeakersCanCreatePointOfOrderForOthers *bool           `json:"list_of_speakers_can_create_point_of_order_for_others"`
	ListOfSpeakersCanSetContributionSelf         *bool           `json:"list_of_speakers_can_set_contribution_self"`
	ListOfSpeakersClosingDisablesPointOfOrder    *bool           `json:"list_of_speakers_closing_disables_point_of_order"`
	ListOfSpeakersCountdownID                    *int            `json:"list_of_speakers_countdown_id"`
	ListOfSpeakersCoupleCountdown                *bool           `json:"list_of_speakers_couple_countdown"`
	ListOfSpeakersDefaultStructureLevelTime      *int            `json:"list_of_speakers_default_structure_level_time"`
	ListOfSpeakersEnableInterposedQuestion       *bool           `json:"list_of_speakers_enable_interposed_question"`
	ListOfSpeakersEnablePointOfOrderCategories   *bool           `json:"list_of_speakers_enable_point_of_order_categories"`
	ListOfSpeakersEnablePointOfOrderSpeakers     *bool           `json:"list_of_speakers_enable_point_of_order_speakers"`
	ListOfSpeakersEnableProContraSpeech          *bool           `json:"list_of_speakers_enable_pro_contra_speech"`
	ListOfSpeakersHideContributionCount          *bool           `json:"list_of_speakers_hide_contribution_count"`
	ListOfSpeakersIDs                            []int           `json:"list_of_speakers_ids"`
	ListOfSpeakersInitiallyClosed                *bool           `json:"list_of_speakers_initially_closed"`
	ListOfSpeakersInterventionTime               *int            `json:"list_of_speakers_intervention_time"`
	ListOfSpeakersPresentUsersOnly               *bool           `json:"list_of_speakers_present_users_only"`
	ListOfSpeakersShowAmountOfSpeakersOnSlide    *bool           `json:"list_of_speakers_show_amount_of_speakers_on_slide"`
	ListOfSpeakersShowFirstContribution          *bool           `json:"list_of_speakers_show_first_contribution"`
	ListOfSpeakersSpeakerNoteForEveryone         *bool           `json:"list_of_speakers_speaker_note_for_everyone"`
	Location                                     *string         `json:"location"`
	LockedFromInside                             *bool           `json:"locked_from_inside"`
	LogoPdfBallotPaperID                         *int            `json:"logo_pdf_ballot_paper_id"`
	LogoPdfFooterLID                             *int            `json:"logo_pdf_footer_l_id"`
	LogoPdfFooterRID                             *int            `json:"logo_pdf_footer_r_id"`
	LogoPdfHeaderLID                             *int            `json:"logo_pdf_header_l_id"`
	LogoPdfHeaderRID                             *int            `json:"logo_pdf_header_r_id"`
	LogoProjectorHeaderID                        *int            `json:"logo_projector_header_id"`
	LogoProjectorMainID                          *int            `json:"logo_projector_main_id"`
	LogoWebHeaderID                              *int            `json:"logo_web_header_id"`
	MediafileIDs                                 []int           `json:"mediafile_ids"`
	MeetingMediafileIDs                          []int           `json:"meeting_mediafile_ids"`
	MeetingUserIDs                               []int           `json:"meeting_user_ids"`
	MotionBlockIDs                               []int           `json:"motion_block_ids"`
	MotionCategoryIDs                            []int           `json:"motion_category_ids"`
	MotionChangeRecommendationIDs                []int           `json:"motion_change_recommendation_ids"`
	MotionCommentIDs                             []int           `json:"motion_comment_ids"`
	MotionCommentSectionIDs                      []int           `json:"motion_comment_section_ids"`
	MotionEditorIDs                              []int           `json:"motion_editor_ids"`
	MotionIDs                                    []int           `json:"motion_ids"`
	MotionPollBallotPaperNumber                  *int            `json:"motion_poll_ballot_paper_number"`
	MotionPollBallotPaperSelection               *string         `json:"motion_poll_ballot_paper_selection"`
	MotionPollDefaultBackend                     *string         `json:"motion_poll_default_backend"`
	MotionPollDefaultGroupIDs                    []int           `json:"motion_poll_default_group_ids"`
	MotionPollDefaultMethod                      *string         `json:"motion_poll_default_method"`
	MotionPollDefaultOnehundredPercentBase       *string         `json:"motion_poll_default_onehundred_percent_base"`
	MotionPollDefaultType                        *string         `json:"motion_poll_default_type"`
	MotionStateIDs                               []int           `json:"motion_state_ids"`
	MotionSubmitterIDs                           []int           `json:"motion_submitter_ids"`
	MotionWorkflowIDs                            []int           `json:"motion_workflow_ids"`
	MotionWorkingGroupSpeakerIDs                 []int           `json:"motion_working_group_speaker_ids"`
	MotionsAmendmentsEnabled                     *bool           `json:"motions_amendments_enabled"`
	MotionsAmendmentsInMainList                  *bool           `json:"motions_amendments_in_main_list"`
	MotionsAmendmentsMultipleParagraphs          *bool           `json:"motions_amendments_multiple_paragraphs"`
	MotionsAmendmentsOfAmendments                *bool           `json:"motions_amendments_of_amendments"`
	MotionsAmendmentsPrefix                      *string         `json:"motions_amendments_prefix"`
	MotionsAmendmentsTextMode                    *string         `json:"motions_amendments_text_mode"`
	MotionsBlockSlideColumns                     *int            `json:"motions_block_slide_columns"`
	MotionsCreateEnableAdditionalSubmitterText   *bool           `json:"motions_create_enable_additional_submitter_text"`
	MotionsDefaultAmendmentWorkflowID            int             `json:"motions_default_amendment_workflow_id"`
	MotionsDefaultLineNumbering                  *string         `json:"motions_default_line_numbering"`
	MotionsDefaultSorting                        *string         `json:"motions_default_sorting"`
	MotionsDefaultWorkflowID                     int             `json:"motions_default_workflow_id"`
	MotionsEnableEditor                          *bool           `json:"motions_enable_editor"`
	MotionsEnableReasonOnProjector               *bool           `json:"motions_enable_reason_on_projector"`
	MotionsEnableRecommendationOnProjector       *bool           `json:"motions_enable_recommendation_on_projector"`
	MotionsEnableSideboxOnProjector              *bool           `json:"motions_enable_sidebox_on_projector"`
	MotionsEnableTextOnProjector                 *bool           `json:"motions_enable_text_on_projector"`
	MotionsEnableWorkingGroupSpeaker             *bool           `json:"motions_enable_working_group_speaker"`
	MotionsExportFollowRecommendation            *bool           `json:"motions_export_follow_recommendation"`
	MotionsExportPreamble                        *string         `json:"motions_export_preamble"`
	MotionsExportSubmitterRecommendation         *bool           `json:"motions_export_submitter_recommendation"`
	MotionsExportTitle                           *string         `json:"motions_export_title"`
	MotionsHideMetadataBackground                *bool           `json:"motions_hide_metadata_background"`
	MotionsLineLength                            *int            `json:"motions_line_length"`
	MotionsNumberMinDigits                       *int            `json:"motions_number_min_digits"`
	MotionsNumberType                            *string         `json:"motions_number_type"`
	MotionsNumberWithBlank                       *bool           `json:"motions_number_with_blank"`
	MotionsPreamble                              *string         `json:"motions_preamble"`
	MotionsReasonRequired                        *bool           `json:"motions_reason_required"`
	MotionsRecommendationTextMode                *string         `json:"motions_recommendation_text_mode"`
	MotionsRecommendationsBy                     *string         `json:"motions_recommendations_by"`
	MotionsShowReferringMotions                  *bool           `json:"motions_show_referring_motions"`
	MotionsShowSequentialNumber                  *bool           `json:"motions_show_sequential_number"`
	MotionsSupportersMinAmount                   *int            `json:"motions_supporters_min_amount"`
	Name                                         string          `json:"name"`
	OptionIDs                                    []int           `json:"option_ids"`
	OrganizationTagIDs                           []int           `json:"organization_tag_ids"`
	PersonalNoteIDs                              []int           `json:"personal_note_ids"`
	PointOfOrderCategoryIDs                      []int           `json:"point_of_order_category_ids"`
	PollBallotPaperNumber                        *int            `json:"poll_ballot_paper_number"`
	PollBallotPaperSelection                     *string         `json:"poll_ballot_paper_selection"`
	PollCandidateIDs                             []int           `json:"poll_candidate_ids"`
	PollCandidateListIDs                         []int           `json:"poll_candidate_list_ids"`
	PollCountdownID                              *int            `json:"poll_countdown_id"`
	PollCoupleCountdown                          *bool           `json:"poll_couple_countdown"`
	PollDefaultBackend                           *string         `json:"poll_default_backend"`
	PollDefaultGroupIDs                          []int           `json:"poll_default_group_ids"`
	PollDefaultMethod                            *string         `json:"poll_default_method"`
	PollDefaultOnehundredPercentBase             *string         `json:"poll_default_onehundred_percent_base"`
	PollDefaultType                              *string         `json:"poll_default_type"`
	PollIDs                                      []int           `json:"poll_ids"`
	PollSortPollResultByVotes                    *bool           `json:"poll_sort_poll_result_by_votes"`
	PresentUserIDs                               []int           `json:"present_user_ids"`
	ProjectionIDs                                []int           `json:"projection_ids"`
	ProjectorCountdownDefaultTime                int             `json:"projector_countdown_default_time"`
	ProjectorCountdownIDs                        []int           `json:"projector_countdown_ids"`
	ProjectorCountdownWarningTime                int             `json:"projector_countdown_warning_time"`
	ProjectorIDs                                 []int           `json:"projector_ids"`
	ProjectorMessageIDs                          []int           `json:"projector_message_ids"`
	ReferenceProjectorID                         int             `json:"reference_projector_id"`
	SpeakerIDs                                   []int           `json:"speaker_ids"`
	StartTime                                    *int            `json:"start_time"`
	StructureLevelIDs                            []int           `json:"structure_level_ids"`
	StructureLevelListOfSpeakersIDs              []int           `json:"structure_level_list_of_speakers_ids"`
	TagIDs                                       []int           `json:"tag_ids"`
	TemplateForOrganizationID                    *int            `json:"template_for_organization_id"`
	TopicIDs                                     []int           `json:"topic_ids"`
	TopicPollDefaultGroupIDs                     []int           `json:"topic_poll_default_group_ids"`
	UserIDs                                      []int           `json:"user_ids"`
	UsersAllowSelfSetPresent                     *bool           `json:"users_allow_self_set_present"`
	UsersEmailBody                               *string         `json:"users_email_body"`
	UsersEmailReplyto                            *string         `json:"users_email_replyto"`
	UsersEmailSender                             *string         `json:"users_email_sender"`
	UsersEmailSubject                            *string         `json:"users_email_subject"`
	UsersEnablePresenceView                      *bool           `json:"users_enable_presence_view"`
	UsersEnableVoteDelegations                   *bool           `json:"users_enable_vote_delegations"`
	UsersEnableVoteWeight                        *bool           `json:"users_enable_vote_weight"`
	UsersForbidDelegatorAsSubmitter              *bool           `json:"users_forbid_delegator_as_submitter"`
	UsersForbidDelegatorAsSupporter              *bool           `json:"users_forbid_delegator_as_supporter"`
	UsersForbidDelegatorInListOfSpeakers         *bool           `json:"users_forbid_delegator_in_list_of_speakers"`
	UsersForbidDelegatorToVote                   *bool           `json:"users_forbid_delegator_to_vote"`
	UsersPdfWelcometext                          *string         `json:"users_pdf_welcometext"`
	UsersPdfWelcometitle                         *string         `json:"users_pdf_welcometitle"`
	UsersPdfWlanEncryption                       *string         `json:"users_pdf_wlan_encryption"`
	UsersPdfWlanPassword                         *string         `json:"users_pdf_wlan_password"`
	UsersPdfWlanSsid                             *string         `json:"users_pdf_wlan_ssid"`
	VoteIDs                                      []int           `json:"vote_ids"`
	WelcomeText                                  *string         `json:"welcome_text"`
	WelcomeTitle                                 *string         `json:"welcome_title"`
	loadedRelations                              map[string]struct{}
	isActiveInOrganization                       *Organization
	templateForOrganization                      *Organization
	projectors                                   []*Projector
	logoWebHeader                                *MeetingMediafile
	agendaItems                                  []*AgendaItem
	defaultProjectorListOfSpeakerss              []*Projector
	defaultProjectorMessages                     []*Projector
	topics                                       []*Topic
	referenceProjector                           *Projector
	defaultProjectorPolls                        []*Projector
	listOfSpeakersCountdown                      *ProjectorCountdown
	motionsDefaultWorkflow                       *MotionWorkflow
	logoPdfFooterR                               *MeetingMediafile
	motionComments                               []*MotionComment
	committee                                    *Committee
	assignmentCandidates                         []*AssignmentCandidate
	fontBold                                     *MeetingMediafile
	pollDefaultGroups                            []*Group
	structureLevels                              []*StructureLevel
	defaultProjectorCurrentListOfSpeakerss       []*Projector
	isArchivedInOrganization                     *Organization
	fontMonospace                                *MeetingMediafile
	presentUsers                                 []*User
	motionWorkingGroupSpeakers                   []*MotionWorkingGroupSpeaker
	polls                                        []*Poll
	topicPollDefaultGroups                       []*Group
	motionsDefaultAmendmentWorkflow              *MotionWorkflow
	fontRegular                                  *MeetingMediafile
	chatGroups                                   []*ChatGroup
	defaultProjectorAssignments                  []*Projector
	defaultProjectorAssignmentPolls              []*Projector
	fontProjectorH2                              *MeetingMediafile
	groups                                       []*Group
	defaultProjectorMediafiles                   []*Projector
	defaultGroup                                 *Group
	motionPollDefaultGroups                      []*Group
	options                                      []*Option
	defaultProjectorTopics                       []*Projector
	logoProjectorHeader                          *MeetingMediafile
	motionStates                                 []*MotionState
	defaultProjectorAgendaItemLists              []*Projector
	listOfSpeakerss                              []*ListOfSpeakers
	motionEditors                                []*MotionEditor
	pollCandidateLists                           []*PollCandidateList
	logoProjectorMain                            *MeetingMediafile
	fontProjectorH1                              *MeetingMediafile
	forwardedMotions                             []*Motion
	projectorCountdowns                          []*ProjectorCountdown
	defaultProjectorCountdowns                   []*Projector
	fontChyronSpeakerName                        *MeetingMediafile
	motions                                      []*Motion
	anonymousGroup                               *Group
	fontItalic                                   *MeetingMediafile
	logoPdfHeaderR                               *MeetingMediafile
	motionChangeRecommendations                  []*MotionChangeRecommendation
	motionBlocks                                 []*MotionBlock
	motionCategorys                              []*MotionCategory
	motionCommentSections                        []*MotionCommentSection
	assignments                                  []*Assignment
	defaultProjectorMotions                      []*Projector
	speakers                                     []*Speaker
	personalNotes                                []*PersonalNote
	logoPdfHeaderL                               *MeetingMediafile
	meetingMediafiles                            []*MeetingMediafile
	motionWorkflows                              []*MotionWorkflow
	defaultProjectorAmendments                   []*Projector
	defaultProjectorMotionBlocks                 []*Projector
	organizationTags                             []*OrganizationTag
	mediafiles                                   []*Mediafile
	chatMessages                                 []*ChatMessage
	meetingUsers                                 []*MeetingUser
	pointOfOrderCategorys                        []*PointOfOrderCategory
	projections                                  []*Projection
	adminGroup                                   *Group
	logoPdfBallotPaper                           *MeetingMediafile
	logoPdfFooterL                               *MeetingMediafile
	tags                                         []*Tag
	defaultMeetingForCommittee                   *Committee
	votes                                        []*Vote
	projectorMessages                            []*ProjectorMessage
	allProjections                               []*Projection
	defaultProjectorMotionPolls                  []*Projector
	pollCandidates                               []*PollCandidate
	pollCountdown                                *ProjectorCountdown
	structureLevelListOfSpeakerss                []*StructureLevelListOfSpeakers
	assignmentPollDefaultGroups                  []*Group
	motionSubmitters                             []*MotionSubmitter
	fontBoldItalic                               *MeetingMediafile
}

func (m *Meeting) CollectionName() string {
	return "meeting"
}

func (m *Meeting) IsActiveInOrganization() *Organization {
	if _, ok := m.loadedRelations["is_active_in_organization_id"]; !ok {
		log.Panic().Msg("Tried to access IsActiveInOrganization relation of Meeting which was not loaded.")
	}

	return m.isActiveInOrganization
}

func (m *Meeting) TemplateForOrganization() *Organization {
	if _, ok := m.loadedRelations["template_for_organization_id"]; !ok {
		log.Panic().Msg("Tried to access TemplateForOrganization relation of Meeting which was not loaded.")
	}

	return m.templateForOrganization
}

func (m *Meeting) Projectors() []*Projector {
	if _, ok := m.loadedRelations["projector_ids"]; !ok {
		log.Panic().Msg("Tried to access Projectors relation of Meeting which was not loaded.")
	}

	return m.projectors
}

func (m *Meeting) LogoWebHeader() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_web_header_id"]; !ok {
		log.Panic().Msg("Tried to access LogoWebHeader relation of Meeting which was not loaded.")
	}

	return m.logoWebHeader
}

func (m *Meeting) AgendaItems() []*AgendaItem {
	if _, ok := m.loadedRelations["agenda_item_ids"]; !ok {
		log.Panic().Msg("Tried to access AgendaItems relation of Meeting which was not loaded.")
	}

	return m.agendaItems
}

func (m *Meeting) DefaultProjectorListOfSpeakerss() []*Projector {
	if _, ok := m.loadedRelations["default_projector_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorListOfSpeakerss
}

func (m *Meeting) DefaultProjectorMessages() []*Projector {
	if _, ok := m.loadedRelations["default_projector_message_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMessages relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMessages
}

func (m *Meeting) Topics() []*Topic {
	if _, ok := m.loadedRelations["topic_ids"]; !ok {
		log.Panic().Msg("Tried to access Topics relation of Meeting which was not loaded.")
	}

	return m.topics
}

func (m *Meeting) ReferenceProjector() Projector {
	if _, ok := m.loadedRelations["reference_projector_id"]; !ok {
		log.Panic().Msg("Tried to access ReferenceProjector relation of Meeting which was not loaded.")
	}

	return *m.referenceProjector
}

func (m *Meeting) DefaultProjectorPolls() []*Projector {
	if _, ok := m.loadedRelations["default_projector_poll_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorPolls relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorPolls
}

func (m *Meeting) ListOfSpeakersCountdown() *ProjectorCountdown {
	if _, ok := m.loadedRelations["list_of_speakers_countdown_id"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakersCountdown relation of Meeting which was not loaded.")
	}

	return m.listOfSpeakersCountdown
}

func (m *Meeting) MotionsDefaultWorkflow() MotionWorkflow {
	if _, ok := m.loadedRelations["motions_default_workflow_id"]; !ok {
		log.Panic().Msg("Tried to access MotionsDefaultWorkflow relation of Meeting which was not loaded.")
	}

	return *m.motionsDefaultWorkflow
}

func (m *Meeting) LogoPdfFooterR() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_footer_r_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfFooterR relation of Meeting which was not loaded.")
	}

	return m.logoPdfFooterR
}

func (m *Meeting) MotionComments() []*MotionComment {
	if _, ok := m.loadedRelations["motion_comment_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionComments relation of Meeting which was not loaded.")
	}

	return m.motionComments
}

func (m *Meeting) Committee() Committee {
	if _, ok := m.loadedRelations["committee_id"]; !ok {
		log.Panic().Msg("Tried to access Committee relation of Meeting which was not loaded.")
	}

	return *m.committee
}

func (m *Meeting) AssignmentCandidates() []*AssignmentCandidate {
	if _, ok := m.loadedRelations["assignment_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access AssignmentCandidates relation of Meeting which was not loaded.")
	}

	return m.assignmentCandidates
}

func (m *Meeting) FontBold() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_bold_id"]; !ok {
		log.Panic().Msg("Tried to access FontBold relation of Meeting which was not loaded.")
	}

	return m.fontBold
}

func (m *Meeting) PollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access PollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.pollDefaultGroups
}

func (m *Meeting) StructureLevels() []*StructureLevel {
	if _, ok := m.loadedRelations["structure_level_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevels relation of Meeting which was not loaded.")
	}

	return m.structureLevels
}

func (m *Meeting) DefaultProjectorCurrentListOfSpeakerss() []*Projector {
	if _, ok := m.loadedRelations["default_projector_current_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorCurrentListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorCurrentListOfSpeakerss
}

func (m *Meeting) IsArchivedInOrganization() *Organization {
	if _, ok := m.loadedRelations["is_archived_in_organization_id"]; !ok {
		log.Panic().Msg("Tried to access IsArchivedInOrganization relation of Meeting which was not loaded.")
	}

	return m.isArchivedInOrganization
}

func (m *Meeting) FontMonospace() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_monospace_id"]; !ok {
		log.Panic().Msg("Tried to access FontMonospace relation of Meeting which was not loaded.")
	}

	return m.fontMonospace
}

func (m *Meeting) PresentUsers() []*User {
	if _, ok := m.loadedRelations["present_user_ids"]; !ok {
		log.Panic().Msg("Tried to access PresentUsers relation of Meeting which was not loaded.")
	}

	return m.presentUsers
}

func (m *Meeting) MotionWorkingGroupSpeakers() []*MotionWorkingGroupSpeaker {
	if _, ok := m.loadedRelations["motion_working_group_speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionWorkingGroupSpeakers relation of Meeting which was not loaded.")
	}

	return m.motionWorkingGroupSpeakers
}

func (m *Meeting) Polls() []*Poll {
	if _, ok := m.loadedRelations["poll_ids"]; !ok {
		log.Panic().Msg("Tried to access Polls relation of Meeting which was not loaded.")
	}

	return m.polls
}

func (m *Meeting) TopicPollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["topic_poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access TopicPollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.topicPollDefaultGroups
}

func (m *Meeting) MotionsDefaultAmendmentWorkflow() MotionWorkflow {
	if _, ok := m.loadedRelations["motions_default_amendment_workflow_id"]; !ok {
		log.Panic().Msg("Tried to access MotionsDefaultAmendmentWorkflow relation of Meeting which was not loaded.")
	}

	return *m.motionsDefaultAmendmentWorkflow
}

func (m *Meeting) FontRegular() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_regular_id"]; !ok {
		log.Panic().Msg("Tried to access FontRegular relation of Meeting which was not loaded.")
	}

	return m.fontRegular
}

func (m *Meeting) ChatGroups() []*ChatGroup {
	if _, ok := m.loadedRelations["chat_group_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatGroups relation of Meeting which was not loaded.")
	}

	return m.chatGroups
}

func (m *Meeting) DefaultProjectorAssignments() []*Projector {
	if _, ok := m.loadedRelations["default_projector_assignment_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAssignments relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAssignments
}

func (m *Meeting) DefaultProjectorAssignmentPolls() []*Projector {
	if _, ok := m.loadedRelations["default_projector_assignment_poll_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAssignmentPolls relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAssignmentPolls
}

func (m *Meeting) FontProjectorH2() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_projector_h2_id"]; !ok {
		log.Panic().Msg("Tried to access FontProjectorH2 relation of Meeting which was not loaded.")
	}

	return m.fontProjectorH2
}

func (m *Meeting) Groups() []*Group {
	if _, ok := m.loadedRelations["group_ids"]; !ok {
		log.Panic().Msg("Tried to access Groups relation of Meeting which was not loaded.")
	}

	return m.groups
}

func (m *Meeting) DefaultProjectorMediafiles() []*Projector {
	if _, ok := m.loadedRelations["default_projector_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMediafiles relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMediafiles
}

func (m *Meeting) DefaultGroup() Group {
	if _, ok := m.loadedRelations["default_group_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultGroup relation of Meeting which was not loaded.")
	}

	return *m.defaultGroup
}

func (m *Meeting) MotionPollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["motion_poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionPollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.motionPollDefaultGroups
}

func (m *Meeting) Options() []*Option {
	if _, ok := m.loadedRelations["option_ids"]; !ok {
		log.Panic().Msg("Tried to access Options relation of Meeting which was not loaded.")
	}

	return m.options
}

func (m *Meeting) DefaultProjectorTopics() []*Projector {
	if _, ok := m.loadedRelations["default_projector_topic_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorTopics relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorTopics
}

func (m *Meeting) LogoProjectorHeader() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_projector_header_id"]; !ok {
		log.Panic().Msg("Tried to access LogoProjectorHeader relation of Meeting which was not loaded.")
	}

	return m.logoProjectorHeader
}

func (m *Meeting) MotionStates() []*MotionState {
	if _, ok := m.loadedRelations["motion_state_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionStates relation of Meeting which was not loaded.")
	}

	return m.motionStates
}

func (m *Meeting) DefaultProjectorAgendaItemLists() []*Projector {
	if _, ok := m.loadedRelations["default_projector_agenda_item_list_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAgendaItemLists relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAgendaItemLists
}

func (m *Meeting) ListOfSpeakerss() []*ListOfSpeakers {
	if _, ok := m.loadedRelations["list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access ListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.listOfSpeakerss
}

func (m *Meeting) MotionEditors() []*MotionEditor {
	if _, ok := m.loadedRelations["motion_editor_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionEditors relation of Meeting which was not loaded.")
	}

	return m.motionEditors
}

func (m *Meeting) PollCandidateLists() []*PollCandidateList {
	if _, ok := m.loadedRelations["poll_candidate_list_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidateLists relation of Meeting which was not loaded.")
	}

	return m.pollCandidateLists
}

func (m *Meeting) LogoProjectorMain() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_projector_main_id"]; !ok {
		log.Panic().Msg("Tried to access LogoProjectorMain relation of Meeting which was not loaded.")
	}

	return m.logoProjectorMain
}

func (m *Meeting) FontProjectorH1() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_projector_h1_id"]; !ok {
		log.Panic().Msg("Tried to access FontProjectorH1 relation of Meeting which was not loaded.")
	}

	return m.fontProjectorH1
}

func (m *Meeting) ForwardedMotions() []*Motion {
	if _, ok := m.loadedRelations["forwarded_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access ForwardedMotions relation of Meeting which was not loaded.")
	}

	return m.forwardedMotions
}

func (m *Meeting) ProjectorCountdowns() []*ProjectorCountdown {
	if _, ok := m.loadedRelations["projector_countdown_ids"]; !ok {
		log.Panic().Msg("Tried to access ProjectorCountdowns relation of Meeting which was not loaded.")
	}

	return m.projectorCountdowns
}

func (m *Meeting) DefaultProjectorCountdowns() []*Projector {
	if _, ok := m.loadedRelations["default_projector_countdown_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorCountdowns relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorCountdowns
}

func (m *Meeting) FontChyronSpeakerName() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_chyron_speaker_name_id"]; !ok {
		log.Panic().Msg("Tried to access FontChyronSpeakerName relation of Meeting which was not loaded.")
	}

	return m.fontChyronSpeakerName
}

func (m *Meeting) Motions() []*Motion {
	if _, ok := m.loadedRelations["motion_ids"]; !ok {
		log.Panic().Msg("Tried to access Motions relation of Meeting which was not loaded.")
	}

	return m.motions
}

func (m *Meeting) AnonymousGroup() *Group {
	if _, ok := m.loadedRelations["anonymous_group_id"]; !ok {
		log.Panic().Msg("Tried to access AnonymousGroup relation of Meeting which was not loaded.")
	}

	return m.anonymousGroup
}

func (m *Meeting) FontItalic() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_italic_id"]; !ok {
		log.Panic().Msg("Tried to access FontItalic relation of Meeting which was not loaded.")
	}

	return m.fontItalic
}

func (m *Meeting) LogoPdfHeaderR() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_header_r_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfHeaderR relation of Meeting which was not loaded.")
	}

	return m.logoPdfHeaderR
}

func (m *Meeting) MotionChangeRecommendations() []*MotionChangeRecommendation {
	if _, ok := m.loadedRelations["motion_change_recommendation_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionChangeRecommendations relation of Meeting which was not loaded.")
	}

	return m.motionChangeRecommendations
}

func (m *Meeting) MotionBlocks() []*MotionBlock {
	if _, ok := m.loadedRelations["motion_block_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionBlocks relation of Meeting which was not loaded.")
	}

	return m.motionBlocks
}

func (m *Meeting) MotionCategorys() []*MotionCategory {
	if _, ok := m.loadedRelations["motion_category_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionCategorys relation of Meeting which was not loaded.")
	}

	return m.motionCategorys
}

func (m *Meeting) MotionCommentSections() []*MotionCommentSection {
	if _, ok := m.loadedRelations["motion_comment_section_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionCommentSections relation of Meeting which was not loaded.")
	}

	return m.motionCommentSections
}

func (m *Meeting) Assignments() []*Assignment {
	if _, ok := m.loadedRelations["assignment_ids"]; !ok {
		log.Panic().Msg("Tried to access Assignments relation of Meeting which was not loaded.")
	}

	return m.assignments
}

func (m *Meeting) DefaultProjectorMotions() []*Projector {
	if _, ok := m.loadedRelations["default_projector_motion_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMotions relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMotions
}

func (m *Meeting) Speakers() []*Speaker {
	if _, ok := m.loadedRelations["speaker_ids"]; !ok {
		log.Panic().Msg("Tried to access Speakers relation of Meeting which was not loaded.")
	}

	return m.speakers
}

func (m *Meeting) PersonalNotes() []*PersonalNote {
	if _, ok := m.loadedRelations["personal_note_ids"]; !ok {
		log.Panic().Msg("Tried to access PersonalNotes relation of Meeting which was not loaded.")
	}

	return m.personalNotes
}

func (m *Meeting) LogoPdfHeaderL() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_header_l_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfHeaderL relation of Meeting which was not loaded.")
	}

	return m.logoPdfHeaderL
}

func (m *Meeting) MeetingMediafiles() []*MeetingMediafile {
	if _, ok := m.loadedRelations["meeting_mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingMediafiles relation of Meeting which was not loaded.")
	}

	return m.meetingMediafiles
}

func (m *Meeting) MotionWorkflows() []*MotionWorkflow {
	if _, ok := m.loadedRelations["motion_workflow_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionWorkflows relation of Meeting which was not loaded.")
	}

	return m.motionWorkflows
}

func (m *Meeting) DefaultProjectorAmendments() []*Projector {
	if _, ok := m.loadedRelations["default_projector_amendment_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorAmendments relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorAmendments
}

func (m *Meeting) DefaultProjectorMotionBlocks() []*Projector {
	if _, ok := m.loadedRelations["default_projector_motion_block_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMotionBlocks relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMotionBlocks
}

func (m *Meeting) OrganizationTags() []*OrganizationTag {
	if _, ok := m.loadedRelations["organization_tag_ids"]; !ok {
		log.Panic().Msg("Tried to access OrganizationTags relation of Meeting which was not loaded.")
	}

	return m.organizationTags
}

func (m *Meeting) Mediafiles() []*Mediafile {
	if _, ok := m.loadedRelations["mediafile_ids"]; !ok {
		log.Panic().Msg("Tried to access Mediafiles relation of Meeting which was not loaded.")
	}

	return m.mediafiles
}

func (m *Meeting) ChatMessages() []*ChatMessage {
	if _, ok := m.loadedRelations["chat_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ChatMessages relation of Meeting which was not loaded.")
	}

	return m.chatMessages
}

func (m *Meeting) MeetingUsers() []*MeetingUser {
	if _, ok := m.loadedRelations["meeting_user_ids"]; !ok {
		log.Panic().Msg("Tried to access MeetingUsers relation of Meeting which was not loaded.")
	}

	return m.meetingUsers
}

func (m *Meeting) PointOfOrderCategorys() []*PointOfOrderCategory {
	if _, ok := m.loadedRelations["point_of_order_category_ids"]; !ok {
		log.Panic().Msg("Tried to access PointOfOrderCategorys relation of Meeting which was not loaded.")
	}

	return m.pointOfOrderCategorys
}

func (m *Meeting) Projections() []*Projection {
	if _, ok := m.loadedRelations["projection_ids"]; !ok {
		log.Panic().Msg("Tried to access Projections relation of Meeting which was not loaded.")
	}

	return m.projections
}

func (m *Meeting) AdminGroup() *Group {
	if _, ok := m.loadedRelations["admin_group_id"]; !ok {
		log.Panic().Msg("Tried to access AdminGroup relation of Meeting which was not loaded.")
	}

	return m.adminGroup
}

func (m *Meeting) LogoPdfBallotPaper() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_ballot_paper_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfBallotPaper relation of Meeting which was not loaded.")
	}

	return m.logoPdfBallotPaper
}

func (m *Meeting) LogoPdfFooterL() *MeetingMediafile {
	if _, ok := m.loadedRelations["logo_pdf_footer_l_id"]; !ok {
		log.Panic().Msg("Tried to access LogoPdfFooterL relation of Meeting which was not loaded.")
	}

	return m.logoPdfFooterL
}

func (m *Meeting) Tags() []*Tag {
	if _, ok := m.loadedRelations["tag_ids"]; !ok {
		log.Panic().Msg("Tried to access Tags relation of Meeting which was not loaded.")
	}

	return m.tags
}

func (m *Meeting) DefaultMeetingForCommittee() *Committee {
	if _, ok := m.loadedRelations["default_meeting_for_committee_id"]; !ok {
		log.Panic().Msg("Tried to access DefaultMeetingForCommittee relation of Meeting which was not loaded.")
	}

	return m.defaultMeetingForCommittee
}

func (m *Meeting) Votes() []*Vote {
	if _, ok := m.loadedRelations["vote_ids"]; !ok {
		log.Panic().Msg("Tried to access Votes relation of Meeting which was not loaded.")
	}

	return m.votes
}

func (m *Meeting) ProjectorMessages() []*ProjectorMessage {
	if _, ok := m.loadedRelations["projector_message_ids"]; !ok {
		log.Panic().Msg("Tried to access ProjectorMessages relation of Meeting which was not loaded.")
	}

	return m.projectorMessages
}

func (m *Meeting) AllProjections() []*Projection {
	if _, ok := m.loadedRelations["all_projection_ids"]; !ok {
		log.Panic().Msg("Tried to access AllProjections relation of Meeting which was not loaded.")
	}

	return m.allProjections
}

func (m *Meeting) DefaultProjectorMotionPolls() []*Projector {
	if _, ok := m.loadedRelations["default_projector_motion_poll_ids"]; !ok {
		log.Panic().Msg("Tried to access DefaultProjectorMotionPolls relation of Meeting which was not loaded.")
	}

	return m.defaultProjectorMotionPolls
}

func (m *Meeting) PollCandidates() []*PollCandidate {
	if _, ok := m.loadedRelations["poll_candidate_ids"]; !ok {
		log.Panic().Msg("Tried to access PollCandidates relation of Meeting which was not loaded.")
	}

	return m.pollCandidates
}

func (m *Meeting) PollCountdown() *ProjectorCountdown {
	if _, ok := m.loadedRelations["poll_countdown_id"]; !ok {
		log.Panic().Msg("Tried to access PollCountdown relation of Meeting which was not loaded.")
	}

	return m.pollCountdown
}

func (m *Meeting) StructureLevelListOfSpeakerss() []*StructureLevelListOfSpeakers {
	if _, ok := m.loadedRelations["structure_level_list_of_speakers_ids"]; !ok {
		log.Panic().Msg("Tried to access StructureLevelListOfSpeakerss relation of Meeting which was not loaded.")
	}

	return m.structureLevelListOfSpeakerss
}

func (m *Meeting) AssignmentPollDefaultGroups() []*Group {
	if _, ok := m.loadedRelations["assignment_poll_default_group_ids"]; !ok {
		log.Panic().Msg("Tried to access AssignmentPollDefaultGroups relation of Meeting which was not loaded.")
	}

	return m.assignmentPollDefaultGroups
}

func (m *Meeting) MotionSubmitters() []*MotionSubmitter {
	if _, ok := m.loadedRelations["motion_submitter_ids"]; !ok {
		log.Panic().Msg("Tried to access MotionSubmitters relation of Meeting which was not loaded.")
	}

	return m.motionSubmitters
}

func (m *Meeting) FontBoldItalic() *MeetingMediafile {
	if _, ok := m.loadedRelations["font_bold_italic_id"]; !ok {
		log.Panic().Msg("Tried to access FontBoldItalic relation of Meeting which was not loaded.")
	}

	return m.fontBoldItalic
}

func (m *Meeting) SetRelated(field string, content interface{}) {
	if content != nil {
		switch field {
		case "is_active_in_organization_id":
			m.isActiveInOrganization = content.(*Organization)
		case "template_for_organization_id":
			m.templateForOrganization = content.(*Organization)
		case "projector_ids":
			m.projectors = content.([]*Projector)
		case "logo_web_header_id":
			m.logoWebHeader = content.(*MeetingMediafile)
		case "agenda_item_ids":
			m.agendaItems = content.([]*AgendaItem)
		case "default_projector_list_of_speakers_ids":
			m.defaultProjectorListOfSpeakerss = content.([]*Projector)
		case "default_projector_message_ids":
			m.defaultProjectorMessages = content.([]*Projector)
		case "topic_ids":
			m.topics = content.([]*Topic)
		case "reference_projector_id":
			m.referenceProjector = content.(*Projector)
		case "default_projector_poll_ids":
			m.defaultProjectorPolls = content.([]*Projector)
		case "list_of_speakers_countdown_id":
			m.listOfSpeakersCountdown = content.(*ProjectorCountdown)
		case "motions_default_workflow_id":
			m.motionsDefaultWorkflow = content.(*MotionWorkflow)
		case "logo_pdf_footer_r_id":
			m.logoPdfFooterR = content.(*MeetingMediafile)
		case "motion_comment_ids":
			m.motionComments = content.([]*MotionComment)
		case "committee_id":
			m.committee = content.(*Committee)
		case "assignment_candidate_ids":
			m.assignmentCandidates = content.([]*AssignmentCandidate)
		case "font_bold_id":
			m.fontBold = content.(*MeetingMediafile)
		case "poll_default_group_ids":
			m.pollDefaultGroups = content.([]*Group)
		case "structure_level_ids":
			m.structureLevels = content.([]*StructureLevel)
		case "default_projector_current_list_of_speakers_ids":
			m.defaultProjectorCurrentListOfSpeakerss = content.([]*Projector)
		case "is_archived_in_organization_id":
			m.isArchivedInOrganization = content.(*Organization)
		case "font_monospace_id":
			m.fontMonospace = content.(*MeetingMediafile)
		case "present_user_ids":
			m.presentUsers = content.([]*User)
		case "motion_working_group_speaker_ids":
			m.motionWorkingGroupSpeakers = content.([]*MotionWorkingGroupSpeaker)
		case "poll_ids":
			m.polls = content.([]*Poll)
		case "topic_poll_default_group_ids":
			m.topicPollDefaultGroups = content.([]*Group)
		case "motions_default_amendment_workflow_id":
			m.motionsDefaultAmendmentWorkflow = content.(*MotionWorkflow)
		case "font_regular_id":
			m.fontRegular = content.(*MeetingMediafile)
		case "chat_group_ids":
			m.chatGroups = content.([]*ChatGroup)
		case "default_projector_assignment_ids":
			m.defaultProjectorAssignments = content.([]*Projector)
		case "default_projector_assignment_poll_ids":
			m.defaultProjectorAssignmentPolls = content.([]*Projector)
		case "font_projector_h2_id":
			m.fontProjectorH2 = content.(*MeetingMediafile)
		case "group_ids":
			m.groups = content.([]*Group)
		case "default_projector_mediafile_ids":
			m.defaultProjectorMediafiles = content.([]*Projector)
		case "default_group_id":
			m.defaultGroup = content.(*Group)
		case "motion_poll_default_group_ids":
			m.motionPollDefaultGroups = content.([]*Group)
		case "option_ids":
			m.options = content.([]*Option)
		case "default_projector_topic_ids":
			m.defaultProjectorTopics = content.([]*Projector)
		case "logo_projector_header_id":
			m.logoProjectorHeader = content.(*MeetingMediafile)
		case "motion_state_ids":
			m.motionStates = content.([]*MotionState)
		case "default_projector_agenda_item_list_ids":
			m.defaultProjectorAgendaItemLists = content.([]*Projector)
		case "list_of_speakers_ids":
			m.listOfSpeakerss = content.([]*ListOfSpeakers)
		case "motion_editor_ids":
			m.motionEditors = content.([]*MotionEditor)
		case "poll_candidate_list_ids":
			m.pollCandidateLists = content.([]*PollCandidateList)
		case "logo_projector_main_id":
			m.logoProjectorMain = content.(*MeetingMediafile)
		case "font_projector_h1_id":
			m.fontProjectorH1 = content.(*MeetingMediafile)
		case "forwarded_motion_ids":
			m.forwardedMotions = content.([]*Motion)
		case "projector_countdown_ids":
			m.projectorCountdowns = content.([]*ProjectorCountdown)
		case "default_projector_countdown_ids":
			m.defaultProjectorCountdowns = content.([]*Projector)
		case "font_chyron_speaker_name_id":
			m.fontChyronSpeakerName = content.(*MeetingMediafile)
		case "motion_ids":
			m.motions = content.([]*Motion)
		case "anonymous_group_id":
			m.anonymousGroup = content.(*Group)
		case "font_italic_id":
			m.fontItalic = content.(*MeetingMediafile)
		case "logo_pdf_header_r_id":
			m.logoPdfHeaderR = content.(*MeetingMediafile)
		case "motion_change_recommendation_ids":
			m.motionChangeRecommendations = content.([]*MotionChangeRecommendation)
		case "motion_block_ids":
			m.motionBlocks = content.([]*MotionBlock)
		case "motion_category_ids":
			m.motionCategorys = content.([]*MotionCategory)
		case "motion_comment_section_ids":
			m.motionCommentSections = content.([]*MotionCommentSection)
		case "assignment_ids":
			m.assignments = content.([]*Assignment)
		case "default_projector_motion_ids":
			m.defaultProjectorMotions = content.([]*Projector)
		case "speaker_ids":
			m.speakers = content.([]*Speaker)
		case "personal_note_ids":
			m.personalNotes = content.([]*PersonalNote)
		case "logo_pdf_header_l_id":
			m.logoPdfHeaderL = content.(*MeetingMediafile)
		case "meeting_mediafile_ids":
			m.meetingMediafiles = content.([]*MeetingMediafile)
		case "motion_workflow_ids":
			m.motionWorkflows = content.([]*MotionWorkflow)
		case "default_projector_amendment_ids":
			m.defaultProjectorAmendments = content.([]*Projector)
		case "default_projector_motion_block_ids":
			m.defaultProjectorMotionBlocks = content.([]*Projector)
		case "organization_tag_ids":
			m.organizationTags = content.([]*OrganizationTag)
		case "mediafile_ids":
			m.mediafiles = content.([]*Mediafile)
		case "chat_message_ids":
			m.chatMessages = content.([]*ChatMessage)
		case "meeting_user_ids":
			m.meetingUsers = content.([]*MeetingUser)
		case "point_of_order_category_ids":
			m.pointOfOrderCategorys = content.([]*PointOfOrderCategory)
		case "projection_ids":
			m.projections = content.([]*Projection)
		case "admin_group_id":
			m.adminGroup = content.(*Group)
		case "logo_pdf_ballot_paper_id":
			m.logoPdfBallotPaper = content.(*MeetingMediafile)
		case "logo_pdf_footer_l_id":
			m.logoPdfFooterL = content.(*MeetingMediafile)
		case "tag_ids":
			m.tags = content.([]*Tag)
		case "default_meeting_for_committee_id":
			m.defaultMeetingForCommittee = content.(*Committee)
		case "vote_ids":
			m.votes = content.([]*Vote)
		case "projector_message_ids":
			m.projectorMessages = content.([]*ProjectorMessage)
		case "all_projection_ids":
			m.allProjections = content.([]*Projection)
		case "default_projector_motion_poll_ids":
			m.defaultProjectorMotionPolls = content.([]*Projector)
		case "poll_candidate_ids":
			m.pollCandidates = content.([]*PollCandidate)
		case "poll_countdown_id":
			m.pollCountdown = content.(*ProjectorCountdown)
		case "structure_level_list_of_speakers_ids":
			m.structureLevelListOfSpeakerss = content.([]*StructureLevelListOfSpeakers)
		case "assignment_poll_default_group_ids":
			m.assignmentPollDefaultGroups = content.([]*Group)
		case "motion_submitter_ids":
			m.motionSubmitters = content.([]*MotionSubmitter)
		case "font_bold_italic_id":
			m.fontBoldItalic = content.(*MeetingMediafile)
		default:
			return
		}
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
}

func (m *Meeting) SetRelatedJSON(field string, content []byte) (*RelatedModelsAccessor, error) {
	var result *RelatedModelsAccessor
	switch field {
	case "is_active_in_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.isActiveInOrganization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "template_for_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.templateForOrganization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "projector_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projectors = append(m.projectors, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "logo_web_header_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoWebHeader = &entry

		result = entry.GetRelatedModelsAccessor()
	case "agenda_item_ids":
		var entry AgendaItem
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.agendaItems = append(m.agendaItems, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_list_of_speakers_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorListOfSpeakerss = append(m.defaultProjectorListOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_message_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMessages = append(m.defaultProjectorMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "topic_ids":
		var entry Topic
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.topics = append(m.topics, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "reference_projector_id":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.referenceProjector = &entry

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_poll_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorPolls = append(m.defaultProjectorPolls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "list_of_speakers_countdown_id":
		var entry ProjectorCountdown
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.listOfSpeakersCountdown = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motions_default_workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionsDefaultWorkflow = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_footer_r_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfFooterR = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_comment_ids":
		var entry MotionComment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionComments = append(m.motionComments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "committee_id":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.committee = &entry

		result = entry.GetRelatedModelsAccessor()
	case "assignment_candidate_ids":
		var entry AssignmentCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignmentCandidates = append(m.assignmentCandidates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "font_bold_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontBold = &entry

		result = entry.GetRelatedModelsAccessor()
	case "poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollDefaultGroups = append(m.pollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "structure_level_ids":
		var entry StructureLevel
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevels = append(m.structureLevels, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_current_list_of_speakers_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorCurrentListOfSpeakerss = append(m.defaultProjectorCurrentListOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "is_archived_in_organization_id":
		var entry Organization
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.isArchivedInOrganization = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_monospace_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontMonospace = &entry

		result = entry.GetRelatedModelsAccessor()
	case "present_user_ids":
		var entry User
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.presentUsers = append(m.presentUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_working_group_speaker_ids":
		var entry MotionWorkingGroupSpeaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionWorkingGroupSpeakers = append(m.motionWorkingGroupSpeakers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_ids":
		var entry Poll
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.polls = append(m.polls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "topic_poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.topicPollDefaultGroups = append(m.topicPollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motions_default_amendment_workflow_id":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionsDefaultAmendmentWorkflow = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_regular_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontRegular = &entry

		result = entry.GetRelatedModelsAccessor()
	case "chat_group_ids":
		var entry ChatGroup
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatGroups = append(m.chatGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_assignment_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAssignments = append(m.defaultProjectorAssignments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_assignment_poll_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAssignmentPolls = append(m.defaultProjectorAssignmentPolls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "font_projector_h2_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontProjectorH2 = &entry

		result = entry.GetRelatedModelsAccessor()
	case "group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.groups = append(m.groups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_mediafile_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMediafiles = append(m.defaultProjectorMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_group_id":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultGroup = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionPollDefaultGroups = append(m.motionPollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "option_ids":
		var entry Option
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.options = append(m.options, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_topic_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorTopics = append(m.defaultProjectorTopics, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "logo_projector_header_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoProjectorHeader = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_state_ids":
		var entry MotionState
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionStates = append(m.motionStates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_agenda_item_list_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAgendaItemLists = append(m.defaultProjectorAgendaItemLists, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "list_of_speakers_ids":
		var entry ListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.listOfSpeakerss = append(m.listOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_editor_ids":
		var entry MotionEditor
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionEditors = append(m.motionEditors, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_list_ids":
		var entry PollCandidateList
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidateLists = append(m.pollCandidateLists, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "logo_projector_main_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoProjectorMain = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_projector_h1_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontProjectorH1 = &entry

		result = entry.GetRelatedModelsAccessor()
	case "forwarded_motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.forwardedMotions = append(m.forwardedMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projector_countdown_ids":
		var entry ProjectorCountdown
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projectorCountdowns = append(m.projectorCountdowns, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_countdown_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorCountdowns = append(m.defaultProjectorCountdowns, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "font_chyron_speaker_name_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontChyronSpeakerName = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_ids":
		var entry Motion
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motions = append(m.motions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "anonymous_group_id":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.anonymousGroup = &entry

		result = entry.GetRelatedModelsAccessor()
	case "font_italic_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontItalic = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_header_r_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfHeaderR = &entry

		result = entry.GetRelatedModelsAccessor()
	case "motion_change_recommendation_ids":
		var entry MotionChangeRecommendation
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionChangeRecommendations = append(m.motionChangeRecommendations, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_block_ids":
		var entry MotionBlock
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionBlocks = append(m.motionBlocks, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_category_ids":
		var entry MotionCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionCategorys = append(m.motionCategorys, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_comment_section_ids":
		var entry MotionCommentSection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionCommentSections = append(m.motionCommentSections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "assignment_ids":
		var entry Assignment
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignments = append(m.assignments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_motion_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMotions = append(m.defaultProjectorMotions, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "speaker_ids":
		var entry Speaker
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.speakers = append(m.speakers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "personal_note_ids":
		var entry PersonalNote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.personalNotes = append(m.personalNotes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_header_l_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfHeaderL = &entry

		result = entry.GetRelatedModelsAccessor()
	case "meeting_mediafile_ids":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingMediafiles = append(m.meetingMediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_workflow_ids":
		var entry MotionWorkflow
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionWorkflows = append(m.motionWorkflows, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_amendment_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorAmendments = append(m.defaultProjectorAmendments, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_motion_block_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMotionBlocks = append(m.defaultProjectorMotionBlocks, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "organization_tag_ids":
		var entry OrganizationTag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.organizationTags = append(m.organizationTags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "mediafile_ids":
		var entry Mediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.mediafiles = append(m.mediafiles, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "chat_message_ids":
		var entry ChatMessage
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.chatMessages = append(m.chatMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "meeting_user_ids":
		var entry MeetingUser
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.meetingUsers = append(m.meetingUsers, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "point_of_order_category_ids":
		var entry PointOfOrderCategory
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pointOfOrderCategorys = append(m.pointOfOrderCategorys, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projections = append(m.projections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "admin_group_id":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.adminGroup = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_ballot_paper_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfBallotPaper = &entry

		result = entry.GetRelatedModelsAccessor()
	case "logo_pdf_footer_l_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.logoPdfFooterL = &entry

		result = entry.GetRelatedModelsAccessor()
	case "tag_ids":
		var entry Tag
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.tags = append(m.tags, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_meeting_for_committee_id":
		var entry Committee
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultMeetingForCommittee = &entry

		result = entry.GetRelatedModelsAccessor()
	case "vote_ids":
		var entry Vote
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.votes = append(m.votes, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "projector_message_ids":
		var entry ProjectorMessage
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.projectorMessages = append(m.projectorMessages, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "all_projection_ids":
		var entry Projection
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.allProjections = append(m.allProjections, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "default_projector_motion_poll_ids":
		var entry Projector
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.defaultProjectorMotionPolls = append(m.defaultProjectorMotionPolls, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_candidate_ids":
		var entry PollCandidate
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCandidates = append(m.pollCandidates, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "poll_countdown_id":
		var entry ProjectorCountdown
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.pollCountdown = &entry

		result = entry.GetRelatedModelsAccessor()
	case "structure_level_list_of_speakers_ids":
		var entry StructureLevelListOfSpeakers
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.structureLevelListOfSpeakerss = append(m.structureLevelListOfSpeakerss, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "assignment_poll_default_group_ids":
		var entry Group
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.assignmentPollDefaultGroups = append(m.assignmentPollDefaultGroups, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "motion_submitter_ids":
		var entry MotionSubmitter
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.motionSubmitters = append(m.motionSubmitters, &entry)

		result = entry.GetRelatedModelsAccessor()
	case "font_bold_italic_id":
		var entry MeetingMediafile
		err := json.Unmarshal(content, &entry)
		if err != nil {
			return nil, err
		}

		m.fontBoldItalic = &entry

		result = entry.GetRelatedModelsAccessor()
	default:
		return nil, fmt.Errorf("set related field json on not existing field")
	}

	if m.loadedRelations == nil {
		m.loadedRelations = map[string]struct{}{}
	}
	m.loadedRelations[field] = struct{}{}
	return result, nil
}

func (m *Meeting) Get(field string) interface{} {
	switch field {
	case "admin_group_id":
		return m.AdminGroupID
	case "agenda_enable_numbering":
		return m.AgendaEnableNumbering
	case "agenda_item_creation":
		return m.AgendaItemCreation
	case "agenda_item_ids":
		return m.AgendaItemIDs
	case "agenda_new_items_default_visibility":
		return m.AgendaNewItemsDefaultVisibility
	case "agenda_number_prefix":
		return m.AgendaNumberPrefix
	case "agenda_numeral_system":
		return m.AgendaNumeralSystem
	case "agenda_show_internal_items_on_projector":
		return m.AgendaShowInternalItemsOnProjector
	case "agenda_show_subtitles":
		return m.AgendaShowSubtitles
	case "agenda_show_topic_navigation_on_detail_view":
		return m.AgendaShowTopicNavigationOnDetailView
	case "all_projection_ids":
		return m.AllProjectionIDs
	case "anonymous_group_id":
		return m.AnonymousGroupID
	case "applause_enable":
		return m.ApplauseEnable
	case "applause_max_amount":
		return m.ApplauseMaxAmount
	case "applause_min_amount":
		return m.ApplauseMinAmount
	case "applause_particle_image_url":
		return m.ApplauseParticleImageUrl
	case "applause_show_level":
		return m.ApplauseShowLevel
	case "applause_timeout":
		return m.ApplauseTimeout
	case "applause_type":
		return m.ApplauseType
	case "assignment_candidate_ids":
		return m.AssignmentCandidateIDs
	case "assignment_ids":
		return m.AssignmentIDs
	case "assignment_poll_add_candidates_to_list_of_speakers":
		return m.AssignmentPollAddCandidatesToListOfSpeakers
	case "assignment_poll_ballot_paper_number":
		return m.AssignmentPollBallotPaperNumber
	case "assignment_poll_ballot_paper_selection":
		return m.AssignmentPollBallotPaperSelection
	case "assignment_poll_default_backend":
		return m.AssignmentPollDefaultBackend
	case "assignment_poll_default_group_ids":
		return m.AssignmentPollDefaultGroupIDs
	case "assignment_poll_default_method":
		return m.AssignmentPollDefaultMethod
	case "assignment_poll_default_onehundred_percent_base":
		return m.AssignmentPollDefaultOnehundredPercentBase
	case "assignment_poll_default_type":
		return m.AssignmentPollDefaultType
	case "assignment_poll_enable_max_votes_per_option":
		return m.AssignmentPollEnableMaxVotesPerOption
	case "assignment_poll_sort_poll_result_by_votes":
		return m.AssignmentPollSortPollResultByVotes
	case "assignments_export_preamble":
		return m.AssignmentsExportPreamble
	case "assignments_export_title":
		return m.AssignmentsExportTitle
	case "chat_group_ids":
		return m.ChatGroupIDs
	case "chat_message_ids":
		return m.ChatMessageIDs
	case "committee_id":
		return m.CommitteeID
	case "conference_auto_connect":
		return m.ConferenceAutoConnect
	case "conference_auto_connect_next_speakers":
		return m.ConferenceAutoConnectNextSpeakers
	case "conference_enable_helpdesk":
		return m.ConferenceEnableHelpdesk
	case "conference_los_restriction":
		return m.ConferenceLosRestriction
	case "conference_open_microphone":
		return m.ConferenceOpenMicrophone
	case "conference_open_video":
		return m.ConferenceOpenVideo
	case "conference_show":
		return m.ConferenceShow
	case "conference_stream_poster_url":
		return m.ConferenceStreamPosterUrl
	case "conference_stream_url":
		return m.ConferenceStreamUrl
	case "custom_translations":
		return m.CustomTranslations
	case "default_group_id":
		return m.DefaultGroupID
	case "default_meeting_for_committee_id":
		return m.DefaultMeetingForCommitteeID
	case "default_projector_agenda_item_list_ids":
		return m.DefaultProjectorAgendaItemListIDs
	case "default_projector_amendment_ids":
		return m.DefaultProjectorAmendmentIDs
	case "default_projector_assignment_ids":
		return m.DefaultProjectorAssignmentIDs
	case "default_projector_assignment_poll_ids":
		return m.DefaultProjectorAssignmentPollIDs
	case "default_projector_countdown_ids":
		return m.DefaultProjectorCountdownIDs
	case "default_projector_current_list_of_speakers_ids":
		return m.DefaultProjectorCurrentListOfSpeakersIDs
	case "default_projector_list_of_speakers_ids":
		return m.DefaultProjectorListOfSpeakersIDs
	case "default_projector_mediafile_ids":
		return m.DefaultProjectorMediafileIDs
	case "default_projector_message_ids":
		return m.DefaultProjectorMessageIDs
	case "default_projector_motion_block_ids":
		return m.DefaultProjectorMotionBlockIDs
	case "default_projector_motion_ids":
		return m.DefaultProjectorMotionIDs
	case "default_projector_motion_poll_ids":
		return m.DefaultProjectorMotionPollIDs
	case "default_projector_poll_ids":
		return m.DefaultProjectorPollIDs
	case "default_projector_topic_ids":
		return m.DefaultProjectorTopicIDs
	case "description":
		return m.Description
	case "enable_anonymous":
		return m.EnableAnonymous
	case "end_time":
		return m.EndTime
	case "export_csv_encoding":
		return m.ExportCsvEncoding
	case "export_csv_separator":
		return m.ExportCsvSeparator
	case "export_pdf_fontsize":
		return m.ExportPdfFontsize
	case "export_pdf_line_height":
		return m.ExportPdfLineHeight
	case "export_pdf_page_margin_bottom":
		return m.ExportPdfPageMarginBottom
	case "export_pdf_page_margin_left":
		return m.ExportPdfPageMarginLeft
	case "export_pdf_page_margin_right":
		return m.ExportPdfPageMarginRight
	case "export_pdf_page_margin_top":
		return m.ExportPdfPageMarginTop
	case "export_pdf_pagenumber_alignment":
		return m.ExportPdfPagenumberAlignment
	case "export_pdf_pagesize":
		return m.ExportPdfPagesize
	case "external_id":
		return m.ExternalID
	case "font_bold_id":
		return m.FontBoldID
	case "font_bold_italic_id":
		return m.FontBoldItalicID
	case "font_chyron_speaker_name_id":
		return m.FontChyronSpeakerNameID
	case "font_italic_id":
		return m.FontItalicID
	case "font_monospace_id":
		return m.FontMonospaceID
	case "font_projector_h1_id":
		return m.FontProjectorH1ID
	case "font_projector_h2_id":
		return m.FontProjectorH2ID
	case "font_regular_id":
		return m.FontRegularID
	case "forwarded_motion_ids":
		return m.ForwardedMotionIDs
	case "group_ids":
		return m.GroupIDs
	case "id":
		return m.ID
	case "imported_at":
		return m.ImportedAt
	case "is_active_in_organization_id":
		return m.IsActiveInOrganizationID
	case "is_archived_in_organization_id":
		return m.IsArchivedInOrganizationID
	case "jitsi_domain":
		return m.JitsiDomain
	case "jitsi_room_name":
		return m.JitsiRoomName
	case "jitsi_room_password":
		return m.JitsiRoomPassword
	case "language":
		return m.Language
	case "list_of_speakers_allow_multiple_speakers":
		return m.ListOfSpeakersAllowMultipleSpeakers
	case "list_of_speakers_amount_last_on_projector":
		return m.ListOfSpeakersAmountLastOnProjector
	case "list_of_speakers_amount_next_on_projector":
		return m.ListOfSpeakersAmountNextOnProjector
	case "list_of_speakers_can_create_point_of_order_for_others":
		return m.ListOfSpeakersCanCreatePointOfOrderForOthers
	case "list_of_speakers_can_set_contribution_self":
		return m.ListOfSpeakersCanSetContributionSelf
	case "list_of_speakers_closing_disables_point_of_order":
		return m.ListOfSpeakersClosingDisablesPointOfOrder
	case "list_of_speakers_countdown_id":
		return m.ListOfSpeakersCountdownID
	case "list_of_speakers_couple_countdown":
		return m.ListOfSpeakersCoupleCountdown
	case "list_of_speakers_default_structure_level_time":
		return m.ListOfSpeakersDefaultStructureLevelTime
	case "list_of_speakers_enable_interposed_question":
		return m.ListOfSpeakersEnableInterposedQuestion
	case "list_of_speakers_enable_point_of_order_categories":
		return m.ListOfSpeakersEnablePointOfOrderCategories
	case "list_of_speakers_enable_point_of_order_speakers":
		return m.ListOfSpeakersEnablePointOfOrderSpeakers
	case "list_of_speakers_enable_pro_contra_speech":
		return m.ListOfSpeakersEnableProContraSpeech
	case "list_of_speakers_hide_contribution_count":
		return m.ListOfSpeakersHideContributionCount
	case "list_of_speakers_ids":
		return m.ListOfSpeakersIDs
	case "list_of_speakers_initially_closed":
		return m.ListOfSpeakersInitiallyClosed
	case "list_of_speakers_intervention_time":
		return m.ListOfSpeakersInterventionTime
	case "list_of_speakers_present_users_only":
		return m.ListOfSpeakersPresentUsersOnly
	case "list_of_speakers_show_amount_of_speakers_on_slide":
		return m.ListOfSpeakersShowAmountOfSpeakersOnSlide
	case "list_of_speakers_show_first_contribution":
		return m.ListOfSpeakersShowFirstContribution
	case "list_of_speakers_speaker_note_for_everyone":
		return m.ListOfSpeakersSpeakerNoteForEveryone
	case "location":
		return m.Location
	case "locked_from_inside":
		return m.LockedFromInside
	case "logo_pdf_ballot_paper_id":
		return m.LogoPdfBallotPaperID
	case "logo_pdf_footer_l_id":
		return m.LogoPdfFooterLID
	case "logo_pdf_footer_r_id":
		return m.LogoPdfFooterRID
	case "logo_pdf_header_l_id":
		return m.LogoPdfHeaderLID
	case "logo_pdf_header_r_id":
		return m.LogoPdfHeaderRID
	case "logo_projector_header_id":
		return m.LogoProjectorHeaderID
	case "logo_projector_main_id":
		return m.LogoProjectorMainID
	case "logo_web_header_id":
		return m.LogoWebHeaderID
	case "mediafile_ids":
		return m.MediafileIDs
	case "meeting_mediafile_ids":
		return m.MeetingMediafileIDs
	case "meeting_user_ids":
		return m.MeetingUserIDs
	case "motion_block_ids":
		return m.MotionBlockIDs
	case "motion_category_ids":
		return m.MotionCategoryIDs
	case "motion_change_recommendation_ids":
		return m.MotionChangeRecommendationIDs
	case "motion_comment_ids":
		return m.MotionCommentIDs
	case "motion_comment_section_ids":
		return m.MotionCommentSectionIDs
	case "motion_editor_ids":
		return m.MotionEditorIDs
	case "motion_ids":
		return m.MotionIDs
	case "motion_poll_ballot_paper_number":
		return m.MotionPollBallotPaperNumber
	case "motion_poll_ballot_paper_selection":
		return m.MotionPollBallotPaperSelection
	case "motion_poll_default_backend":
		return m.MotionPollDefaultBackend
	case "motion_poll_default_group_ids":
		return m.MotionPollDefaultGroupIDs
	case "motion_poll_default_method":
		return m.MotionPollDefaultMethod
	case "motion_poll_default_onehundred_percent_base":
		return m.MotionPollDefaultOnehundredPercentBase
	case "motion_poll_default_type":
		return m.MotionPollDefaultType
	case "motion_state_ids":
		return m.MotionStateIDs
	case "motion_submitter_ids":
		return m.MotionSubmitterIDs
	case "motion_workflow_ids":
		return m.MotionWorkflowIDs
	case "motion_working_group_speaker_ids":
		return m.MotionWorkingGroupSpeakerIDs
	case "motions_amendments_enabled":
		return m.MotionsAmendmentsEnabled
	case "motions_amendments_in_main_list":
		return m.MotionsAmendmentsInMainList
	case "motions_amendments_multiple_paragraphs":
		return m.MotionsAmendmentsMultipleParagraphs
	case "motions_amendments_of_amendments":
		return m.MotionsAmendmentsOfAmendments
	case "motions_amendments_prefix":
		return m.MotionsAmendmentsPrefix
	case "motions_amendments_text_mode":
		return m.MotionsAmendmentsTextMode
	case "motions_block_slide_columns":
		return m.MotionsBlockSlideColumns
	case "motions_create_enable_additional_submitter_text":
		return m.MotionsCreateEnableAdditionalSubmitterText
	case "motions_default_amendment_workflow_id":
		return m.MotionsDefaultAmendmentWorkflowID
	case "motions_default_line_numbering":
		return m.MotionsDefaultLineNumbering
	case "motions_default_sorting":
		return m.MotionsDefaultSorting
	case "motions_default_workflow_id":
		return m.MotionsDefaultWorkflowID
	case "motions_enable_editor":
		return m.MotionsEnableEditor
	case "motions_enable_reason_on_projector":
		return m.MotionsEnableReasonOnProjector
	case "motions_enable_recommendation_on_projector":
		return m.MotionsEnableRecommendationOnProjector
	case "motions_enable_sidebox_on_projector":
		return m.MotionsEnableSideboxOnProjector
	case "motions_enable_text_on_projector":
		return m.MotionsEnableTextOnProjector
	case "motions_enable_working_group_speaker":
		return m.MotionsEnableWorkingGroupSpeaker
	case "motions_export_follow_recommendation":
		return m.MotionsExportFollowRecommendation
	case "motions_export_preamble":
		return m.MotionsExportPreamble
	case "motions_export_submitter_recommendation":
		return m.MotionsExportSubmitterRecommendation
	case "motions_export_title":
		return m.MotionsExportTitle
	case "motions_hide_metadata_background":
		return m.MotionsHideMetadataBackground
	case "motions_line_length":
		return m.MotionsLineLength
	case "motions_number_min_digits":
		return m.MotionsNumberMinDigits
	case "motions_number_type":
		return m.MotionsNumberType
	case "motions_number_with_blank":
		return m.MotionsNumberWithBlank
	case "motions_preamble":
		return m.MotionsPreamble
	case "motions_reason_required":
		return m.MotionsReasonRequired
	case "motions_recommendation_text_mode":
		return m.MotionsRecommendationTextMode
	case "motions_recommendations_by":
		return m.MotionsRecommendationsBy
	case "motions_show_referring_motions":
		return m.MotionsShowReferringMotions
	case "motions_show_sequential_number":
		return m.MotionsShowSequentialNumber
	case "motions_supporters_min_amount":
		return m.MotionsSupportersMinAmount
	case "name":
		return m.Name
	case "option_ids":
		return m.OptionIDs
	case "organization_tag_ids":
		return m.OrganizationTagIDs
	case "personal_note_ids":
		return m.PersonalNoteIDs
	case "point_of_order_category_ids":
		return m.PointOfOrderCategoryIDs
	case "poll_ballot_paper_number":
		return m.PollBallotPaperNumber
	case "poll_ballot_paper_selection":
		return m.PollBallotPaperSelection
	case "poll_candidate_ids":
		return m.PollCandidateIDs
	case "poll_candidate_list_ids":
		return m.PollCandidateListIDs
	case "poll_countdown_id":
		return m.PollCountdownID
	case "poll_couple_countdown":
		return m.PollCoupleCountdown
	case "poll_default_backend":
		return m.PollDefaultBackend
	case "poll_default_group_ids":
		return m.PollDefaultGroupIDs
	case "poll_default_method":
		return m.PollDefaultMethod
	case "poll_default_onehundred_percent_base":
		return m.PollDefaultOnehundredPercentBase
	case "poll_default_type":
		return m.PollDefaultType
	case "poll_ids":
		return m.PollIDs
	case "poll_sort_poll_result_by_votes":
		return m.PollSortPollResultByVotes
	case "present_user_ids":
		return m.PresentUserIDs
	case "projection_ids":
		return m.ProjectionIDs
	case "projector_countdown_default_time":
		return m.ProjectorCountdownDefaultTime
	case "projector_countdown_ids":
		return m.ProjectorCountdownIDs
	case "projector_countdown_warning_time":
		return m.ProjectorCountdownWarningTime
	case "projector_ids":
		return m.ProjectorIDs
	case "projector_message_ids":
		return m.ProjectorMessageIDs
	case "reference_projector_id":
		return m.ReferenceProjectorID
	case "speaker_ids":
		return m.SpeakerIDs
	case "start_time":
		return m.StartTime
	case "structure_level_ids":
		return m.StructureLevelIDs
	case "structure_level_list_of_speakers_ids":
		return m.StructureLevelListOfSpeakersIDs
	case "tag_ids":
		return m.TagIDs
	case "template_for_organization_id":
		return m.TemplateForOrganizationID
	case "topic_ids":
		return m.TopicIDs
	case "topic_poll_default_group_ids":
		return m.TopicPollDefaultGroupIDs
	case "user_ids":
		return m.UserIDs
	case "users_allow_self_set_present":
		return m.UsersAllowSelfSetPresent
	case "users_email_body":
		return m.UsersEmailBody
	case "users_email_replyto":
		return m.UsersEmailReplyto
	case "users_email_sender":
		return m.UsersEmailSender
	case "users_email_subject":
		return m.UsersEmailSubject
	case "users_enable_presence_view":
		return m.UsersEnablePresenceView
	case "users_enable_vote_delegations":
		return m.UsersEnableVoteDelegations
	case "users_enable_vote_weight":
		return m.UsersEnableVoteWeight
	case "users_forbid_delegator_as_submitter":
		return m.UsersForbidDelegatorAsSubmitter
	case "users_forbid_delegator_as_supporter":
		return m.UsersForbidDelegatorAsSupporter
	case "users_forbid_delegator_in_list_of_speakers":
		return m.UsersForbidDelegatorInListOfSpeakers
	case "users_forbid_delegator_to_vote":
		return m.UsersForbidDelegatorToVote
	case "users_pdf_welcometext":
		return m.UsersPdfWelcometext
	case "users_pdf_welcometitle":
		return m.UsersPdfWelcometitle
	case "users_pdf_wlan_encryption":
		return m.UsersPdfWlanEncryption
	case "users_pdf_wlan_password":
		return m.UsersPdfWlanPassword
	case "users_pdf_wlan_ssid":
		return m.UsersPdfWlanSsid
	case "vote_ids":
		return m.VoteIDs
	case "welcome_text":
		return m.WelcomeText
	case "welcome_title":
		return m.WelcomeTitle
	}

	return nil
}

func (m *Meeting) GetFqids(field string) []string {
	switch field {
	case "is_active_in_organization_id":
		if m.IsActiveInOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.IsActiveInOrganizationID)}
		}

	case "template_for_organization_id":
		if m.TemplateForOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.TemplateForOrganizationID)}
		}

	case "projector_ids":
		r := make([]string, len(m.ProjectorIDs))
		for i, id := range m.ProjectorIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "logo_web_header_id":
		if m.LogoWebHeaderID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoWebHeaderID)}
		}

	case "agenda_item_ids":
		r := make([]string, len(m.AgendaItemIDs))
		for i, id := range m.AgendaItemIDs {
			r[i] = "agenda_item/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_list_of_speakers_ids":
		r := make([]string, len(m.DefaultProjectorListOfSpeakersIDs))
		for i, id := range m.DefaultProjectorListOfSpeakersIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_message_ids":
		r := make([]string, len(m.DefaultProjectorMessageIDs))
		for i, id := range m.DefaultProjectorMessageIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "topic_ids":
		r := make([]string, len(m.TopicIDs))
		for i, id := range m.TopicIDs {
			r[i] = "topic/" + strconv.Itoa(id)
		}
		return r

	case "reference_projector_id":
		return []string{"projector/" + strconv.Itoa(m.ReferenceProjectorID)}

	case "default_projector_poll_ids":
		r := make([]string, len(m.DefaultProjectorPollIDs))
		for i, id := range m.DefaultProjectorPollIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "list_of_speakers_countdown_id":
		if m.ListOfSpeakersCountdownID != nil {
			return []string{"projector_countdown/" + strconv.Itoa(*m.ListOfSpeakersCountdownID)}
		}

	case "motions_default_workflow_id":
		return []string{"motion_workflow/" + strconv.Itoa(m.MotionsDefaultWorkflowID)}

	case "logo_pdf_footer_r_id":
		if m.LogoPdfFooterRID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfFooterRID)}
		}

	case "motion_comment_ids":
		r := make([]string, len(m.MotionCommentIDs))
		for i, id := range m.MotionCommentIDs {
			r[i] = "motion_comment/" + strconv.Itoa(id)
		}
		return r

	case "committee_id":
		return []string{"committee/" + strconv.Itoa(m.CommitteeID)}

	case "assignment_candidate_ids":
		r := make([]string, len(m.AssignmentCandidateIDs))
		for i, id := range m.AssignmentCandidateIDs {
			r[i] = "assignment_candidate/" + strconv.Itoa(id)
		}
		return r

	case "font_bold_id":
		if m.FontBoldID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontBoldID)}
		}

	case "poll_default_group_ids":
		r := make([]string, len(m.PollDefaultGroupIDs))
		for i, id := range m.PollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "structure_level_ids":
		r := make([]string, len(m.StructureLevelIDs))
		for i, id := range m.StructureLevelIDs {
			r[i] = "structure_level/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_current_list_of_speakers_ids":
		r := make([]string, len(m.DefaultProjectorCurrentListOfSpeakersIDs))
		for i, id := range m.DefaultProjectorCurrentListOfSpeakersIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "is_archived_in_organization_id":
		if m.IsArchivedInOrganizationID != nil {
			return []string{"organization/" + strconv.Itoa(*m.IsArchivedInOrganizationID)}
		}

	case "font_monospace_id":
		if m.FontMonospaceID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontMonospaceID)}
		}

	case "present_user_ids":
		r := make([]string, len(m.PresentUserIDs))
		for i, id := range m.PresentUserIDs {
			r[i] = "user/" + strconv.Itoa(id)
		}
		return r

	case "motion_working_group_speaker_ids":
		r := make([]string, len(m.MotionWorkingGroupSpeakerIDs))
		for i, id := range m.MotionWorkingGroupSpeakerIDs {
			r[i] = "motion_working_group_speaker/" + strconv.Itoa(id)
		}
		return r

	case "poll_ids":
		r := make([]string, len(m.PollIDs))
		for i, id := range m.PollIDs {
			r[i] = "poll/" + strconv.Itoa(id)
		}
		return r

	case "topic_poll_default_group_ids":
		r := make([]string, len(m.TopicPollDefaultGroupIDs))
		for i, id := range m.TopicPollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "motions_default_amendment_workflow_id":
		return []string{"motion_workflow/" + strconv.Itoa(m.MotionsDefaultAmendmentWorkflowID)}

	case "font_regular_id":
		if m.FontRegularID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontRegularID)}
		}

	case "chat_group_ids":
		r := make([]string, len(m.ChatGroupIDs))
		for i, id := range m.ChatGroupIDs {
			r[i] = "chat_group/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_assignment_ids":
		r := make([]string, len(m.DefaultProjectorAssignmentIDs))
		for i, id := range m.DefaultProjectorAssignmentIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_assignment_poll_ids":
		r := make([]string, len(m.DefaultProjectorAssignmentPollIDs))
		for i, id := range m.DefaultProjectorAssignmentPollIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "font_projector_h2_id":
		if m.FontProjectorH2ID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontProjectorH2ID)}
		}

	case "group_ids":
		r := make([]string, len(m.GroupIDs))
		for i, id := range m.GroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_mediafile_ids":
		r := make([]string, len(m.DefaultProjectorMediafileIDs))
		for i, id := range m.DefaultProjectorMediafileIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_group_id":
		return []string{"group/" + strconv.Itoa(m.DefaultGroupID)}

	case "motion_poll_default_group_ids":
		r := make([]string, len(m.MotionPollDefaultGroupIDs))
		for i, id := range m.MotionPollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "option_ids":
		r := make([]string, len(m.OptionIDs))
		for i, id := range m.OptionIDs {
			r[i] = "option/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_topic_ids":
		r := make([]string, len(m.DefaultProjectorTopicIDs))
		for i, id := range m.DefaultProjectorTopicIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "logo_projector_header_id":
		if m.LogoProjectorHeaderID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoProjectorHeaderID)}
		}

	case "motion_state_ids":
		r := make([]string, len(m.MotionStateIDs))
		for i, id := range m.MotionStateIDs {
			r[i] = "motion_state/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_agenda_item_list_ids":
		r := make([]string, len(m.DefaultProjectorAgendaItemListIDs))
		for i, id := range m.DefaultProjectorAgendaItemListIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "list_of_speakers_ids":
		r := make([]string, len(m.ListOfSpeakersIDs))
		for i, id := range m.ListOfSpeakersIDs {
			r[i] = "list_of_speakers/" + strconv.Itoa(id)
		}
		return r

	case "motion_editor_ids":
		r := make([]string, len(m.MotionEditorIDs))
		for i, id := range m.MotionEditorIDs {
			r[i] = "motion_editor/" + strconv.Itoa(id)
		}
		return r

	case "poll_candidate_list_ids":
		r := make([]string, len(m.PollCandidateListIDs))
		for i, id := range m.PollCandidateListIDs {
			r[i] = "poll_candidate_list/" + strconv.Itoa(id)
		}
		return r

	case "logo_projector_main_id":
		if m.LogoProjectorMainID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoProjectorMainID)}
		}

	case "font_projector_h1_id":
		if m.FontProjectorH1ID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontProjectorH1ID)}
		}

	case "forwarded_motion_ids":
		r := make([]string, len(m.ForwardedMotionIDs))
		for i, id := range m.ForwardedMotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "projector_countdown_ids":
		r := make([]string, len(m.ProjectorCountdownIDs))
		for i, id := range m.ProjectorCountdownIDs {
			r[i] = "projector_countdown/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_countdown_ids":
		r := make([]string, len(m.DefaultProjectorCountdownIDs))
		for i, id := range m.DefaultProjectorCountdownIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "font_chyron_speaker_name_id":
		if m.FontChyronSpeakerNameID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontChyronSpeakerNameID)}
		}

	case "motion_ids":
		r := make([]string, len(m.MotionIDs))
		for i, id := range m.MotionIDs {
			r[i] = "motion/" + strconv.Itoa(id)
		}
		return r

	case "anonymous_group_id":
		if m.AnonymousGroupID != nil {
			return []string{"group/" + strconv.Itoa(*m.AnonymousGroupID)}
		}

	case "font_italic_id":
		if m.FontItalicID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontItalicID)}
		}

	case "logo_pdf_header_r_id":
		if m.LogoPdfHeaderRID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfHeaderRID)}
		}

	case "motion_change_recommendation_ids":
		r := make([]string, len(m.MotionChangeRecommendationIDs))
		for i, id := range m.MotionChangeRecommendationIDs {
			r[i] = "motion_change_recommendation/" + strconv.Itoa(id)
		}
		return r

	case "motion_block_ids":
		r := make([]string, len(m.MotionBlockIDs))
		for i, id := range m.MotionBlockIDs {
			r[i] = "motion_block/" + strconv.Itoa(id)
		}
		return r

	case "motion_category_ids":
		r := make([]string, len(m.MotionCategoryIDs))
		for i, id := range m.MotionCategoryIDs {
			r[i] = "motion_category/" + strconv.Itoa(id)
		}
		return r

	case "motion_comment_section_ids":
		r := make([]string, len(m.MotionCommentSectionIDs))
		for i, id := range m.MotionCommentSectionIDs {
			r[i] = "motion_comment_section/" + strconv.Itoa(id)
		}
		return r

	case "assignment_ids":
		r := make([]string, len(m.AssignmentIDs))
		for i, id := range m.AssignmentIDs {
			r[i] = "assignment/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_motion_ids":
		r := make([]string, len(m.DefaultProjectorMotionIDs))
		for i, id := range m.DefaultProjectorMotionIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "speaker_ids":
		r := make([]string, len(m.SpeakerIDs))
		for i, id := range m.SpeakerIDs {
			r[i] = "speaker/" + strconv.Itoa(id)
		}
		return r

	case "personal_note_ids":
		r := make([]string, len(m.PersonalNoteIDs))
		for i, id := range m.PersonalNoteIDs {
			r[i] = "personal_note/" + strconv.Itoa(id)
		}
		return r

	case "logo_pdf_header_l_id":
		if m.LogoPdfHeaderLID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfHeaderLID)}
		}

	case "meeting_mediafile_ids":
		r := make([]string, len(m.MeetingMediafileIDs))
		for i, id := range m.MeetingMediafileIDs {
			r[i] = "meeting_mediafile/" + strconv.Itoa(id)
		}
		return r

	case "motion_workflow_ids":
		r := make([]string, len(m.MotionWorkflowIDs))
		for i, id := range m.MotionWorkflowIDs {
			r[i] = "motion_workflow/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_amendment_ids":
		r := make([]string, len(m.DefaultProjectorAmendmentIDs))
		for i, id := range m.DefaultProjectorAmendmentIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_motion_block_ids":
		r := make([]string, len(m.DefaultProjectorMotionBlockIDs))
		for i, id := range m.DefaultProjectorMotionBlockIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "organization_tag_ids":
		r := make([]string, len(m.OrganizationTagIDs))
		for i, id := range m.OrganizationTagIDs {
			r[i] = "organization_tag/" + strconv.Itoa(id)
		}
		return r

	case "mediafile_ids":
		r := make([]string, len(m.MediafileIDs))
		for i, id := range m.MediafileIDs {
			r[i] = "mediafile/" + strconv.Itoa(id)
		}
		return r

	case "chat_message_ids":
		r := make([]string, len(m.ChatMessageIDs))
		for i, id := range m.ChatMessageIDs {
			r[i] = "chat_message/" + strconv.Itoa(id)
		}
		return r

	case "meeting_user_ids":
		r := make([]string, len(m.MeetingUserIDs))
		for i, id := range m.MeetingUserIDs {
			r[i] = "meeting_user/" + strconv.Itoa(id)
		}
		return r

	case "point_of_order_category_ids":
		r := make([]string, len(m.PointOfOrderCategoryIDs))
		for i, id := range m.PointOfOrderCategoryIDs {
			r[i] = "point_of_order_category/" + strconv.Itoa(id)
		}
		return r

	case "projection_ids":
		r := make([]string, len(m.ProjectionIDs))
		for i, id := range m.ProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "admin_group_id":
		if m.AdminGroupID != nil {
			return []string{"group/" + strconv.Itoa(*m.AdminGroupID)}
		}

	case "logo_pdf_ballot_paper_id":
		if m.LogoPdfBallotPaperID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfBallotPaperID)}
		}

	case "logo_pdf_footer_l_id":
		if m.LogoPdfFooterLID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.LogoPdfFooterLID)}
		}

	case "tag_ids":
		r := make([]string, len(m.TagIDs))
		for i, id := range m.TagIDs {
			r[i] = "tag/" + strconv.Itoa(id)
		}
		return r

	case "default_meeting_for_committee_id":
		if m.DefaultMeetingForCommitteeID != nil {
			return []string{"committee/" + strconv.Itoa(*m.DefaultMeetingForCommitteeID)}
		}

	case "vote_ids":
		r := make([]string, len(m.VoteIDs))
		for i, id := range m.VoteIDs {
			r[i] = "vote/" + strconv.Itoa(id)
		}
		return r

	case "projector_message_ids":
		r := make([]string, len(m.ProjectorMessageIDs))
		for i, id := range m.ProjectorMessageIDs {
			r[i] = "projector_message/" + strconv.Itoa(id)
		}
		return r

	case "all_projection_ids":
		r := make([]string, len(m.AllProjectionIDs))
		for i, id := range m.AllProjectionIDs {
			r[i] = "projection/" + strconv.Itoa(id)
		}
		return r

	case "default_projector_motion_poll_ids":
		r := make([]string, len(m.DefaultProjectorMotionPollIDs))
		for i, id := range m.DefaultProjectorMotionPollIDs {
			r[i] = "projector/" + strconv.Itoa(id)
		}
		return r

	case "poll_candidate_ids":
		r := make([]string, len(m.PollCandidateIDs))
		for i, id := range m.PollCandidateIDs {
			r[i] = "poll_candidate/" + strconv.Itoa(id)
		}
		return r

	case "poll_countdown_id":
		if m.PollCountdownID != nil {
			return []string{"projector_countdown/" + strconv.Itoa(*m.PollCountdownID)}
		}

	case "structure_level_list_of_speakers_ids":
		r := make([]string, len(m.StructureLevelListOfSpeakersIDs))
		for i, id := range m.StructureLevelListOfSpeakersIDs {
			r[i] = "structure_level_list_of_speakers/" + strconv.Itoa(id)
		}
		return r

	case "assignment_poll_default_group_ids":
		r := make([]string, len(m.AssignmentPollDefaultGroupIDs))
		for i, id := range m.AssignmentPollDefaultGroupIDs {
			r[i] = "group/" + strconv.Itoa(id)
		}
		return r

	case "motion_submitter_ids":
		r := make([]string, len(m.MotionSubmitterIDs))
		for i, id := range m.MotionSubmitterIDs {
			r[i] = "motion_submitter/" + strconv.Itoa(id)
		}
		return r

	case "font_bold_italic_id":
		if m.FontBoldItalicID != nil {
			return []string{"meeting_mediafile/" + strconv.Itoa(*m.FontBoldItalicID)}
		}
	}
	return []string{}
}

func (m *Meeting) Update(data map[string]string) error {
	if val, ok := data["admin_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AdminGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_enable_numbering"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaEnableNumbering)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_creation"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemCreation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_item_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaItemIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_new_items_default_visibility"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNewItemsDefaultVisibility)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_number_prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNumberPrefix)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_numeral_system"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaNumeralSystem)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_internal_items_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowInternalItemsOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_subtitles"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowSubtitles)
		if err != nil {
			return err
		}
	}

	if val, ok := data["agenda_show_topic_navigation_on_detail_view"]; ok {
		err := json.Unmarshal([]byte(val), &m.AgendaShowTopicNavigationOnDetailView)
		if err != nil {
			return err
		}
	}

	if val, ok := data["all_projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AllProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["anonymous_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.AnonymousGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_enable"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseEnable)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_max_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseMaxAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_min_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseMinAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_particle_image_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseParticleImageUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_show_level"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseShowLevel)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_timeout"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseTimeout)
		if err != nil {
			return err
		}
	}

	if val, ok := data["applause_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.ApplauseType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentCandidateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_add_candidates_to_list_of_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollAddCandidatesToListOfSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_enable_max_votes_per_option"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollEnableMaxVotesPerOption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignment_poll_sort_poll_result_by_votes"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentPollSortPollResultByVotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignments_export_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentsExportPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["assignments_export_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.AssignmentsExportTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chat_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["chat_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ChatMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["committee_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.CommitteeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_auto_connect"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceAutoConnect)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_auto_connect_next_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceAutoConnectNextSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_enable_helpdesk"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceEnableHelpdesk)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_los_restriction"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceLosRestriction)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_open_microphone"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceOpenMicrophone)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_open_video"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceOpenVideo)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_show"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceShow)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_stream_poster_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceStreamPosterUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["conference_stream_url"]; ok {
		err := json.Unmarshal([]byte(val), &m.ConferenceStreamUrl)
		if err != nil {
			return err
		}
	}

	if val, ok := data["custom_translations"]; ok {
		err := json.Unmarshal([]byte(val), &m.CustomTranslations)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_group_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultGroupID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_meeting_for_committee_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultMeetingForCommitteeID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_agenda_item_list_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAgendaItemListIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_amendment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAmendmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_assignment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAssignmentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_assignment_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorAssignmentPollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_countdown_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorCountdownIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_current_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorCurrentListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_motion_block_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionBlockIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_motion_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorMotionPollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorPollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["default_projector_topic_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.DefaultProjectorTopicIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["description"]; ok {
		err := json.Unmarshal([]byte(val), &m.Description)
		if err != nil {
			return err
		}
	}

	if val, ok := data["enable_anonymous"]; ok {
		err := json.Unmarshal([]byte(val), &m.EnableAnonymous)
		if err != nil {
			return err
		}
	}

	if val, ok := data["end_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.EndTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_csv_encoding"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportCsvEncoding)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_csv_separator"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportCsvSeparator)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_fontsize"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfFontsize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_line_height"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfLineHeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_bottom"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginBottom)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_left"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginLeft)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_right"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginRight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_page_margin_top"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPageMarginTop)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_pagenumber_alignment"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPagenumberAlignment)
		if err != nil {
			return err
		}
	}

	if val, ok := data["export_pdf_pagesize"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExportPdfPagesize)
		if err != nil {
			return err
		}
	}

	if val, ok := data["external_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ExternalID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_bold_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontBoldID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_bold_italic_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontBoldItalicID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_chyron_speaker_name_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontChyronSpeakerNameID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_italic_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontItalicID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_monospace_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontMonospaceID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_projector_h1_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontProjectorH1ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_projector_h2_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontProjectorH2ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["font_regular_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.FontRegularID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["forwarded_motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ForwardedMotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.GroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["imported_at"]; ok {
		err := json.Unmarshal([]byte(val), &m.ImportedAt)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_active_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsActiveInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["is_archived_in_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.IsArchivedInOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_domain"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiDomain)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_room_name"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiRoomName)
		if err != nil {
			return err
		}
	}

	if val, ok := data["jitsi_room_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.JitsiRoomPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["language"]; ok {
		err := json.Unmarshal([]byte(val), &m.Language)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_allow_multiple_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAllowMultipleSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_amount_last_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAmountLastOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_amount_next_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersAmountNextOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_can_create_point_of_order_for_others"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCanCreatePointOfOrderForOthers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_can_set_contribution_self"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCanSetContributionSelf)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_closing_disables_point_of_order"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersClosingDisablesPointOfOrder)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_countdown_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCountdownID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_couple_countdown"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersCoupleCountdown)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_default_structure_level_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersDefaultStructureLevelTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_interposed_question"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnableInterposedQuestion)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_point_of_order_categories"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnablePointOfOrderCategories)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_point_of_order_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnablePointOfOrderSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_enable_pro_contra_speech"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersEnableProContraSpeech)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_hide_contribution_count"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersHideContributionCount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_initially_closed"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersInitiallyClosed)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_intervention_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersInterventionTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_present_users_only"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersPresentUsersOnly)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_show_amount_of_speakers_on_slide"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersShowAmountOfSpeakersOnSlide)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_show_first_contribution"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersShowFirstContribution)
		if err != nil {
			return err
		}
	}

	if val, ok := data["list_of_speakers_speaker_note_for_everyone"]; ok {
		err := json.Unmarshal([]byte(val), &m.ListOfSpeakersSpeakerNoteForEveryone)
		if err != nil {
			return err
		}
	}

	if val, ok := data["location"]; ok {
		err := json.Unmarshal([]byte(val), &m.Location)
		if err != nil {
			return err
		}
	}

	if val, ok := data["locked_from_inside"]; ok {
		err := json.Unmarshal([]byte(val), &m.LockedFromInside)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_ballot_paper_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfBallotPaperID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_footer_l_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfFooterLID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_footer_r_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfFooterRID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_header_l_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfHeaderLID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_pdf_header_r_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoPdfHeaderRID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_projector_header_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoProjectorHeaderID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_projector_main_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoProjectorMainID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["logo_web_header_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.LogoWebHeaderID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_mediafile_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingMediafileIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["meeting_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MeetingUserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_block_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionBlockIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_category_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCategoryIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_change_recommendation_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionChangeRecommendationIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_comment_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCommentIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_comment_section_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionCommentSectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_editor_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionEditorIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionPollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_state_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionStateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_submitter_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionSubmitterIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_workflow_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkflowIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motion_working_group_speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionWorkingGroupSpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_enabled"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsEnabled)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_in_main_list"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsInMainList)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_multiple_paragraphs"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsMultipleParagraphs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_of_amendments"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsOfAmendments)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_prefix"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsPrefix)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_amendments_text_mode"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsAmendmentsTextMode)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_block_slide_columns"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsBlockSlideColumns)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_create_enable_additional_submitter_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsCreateEnableAdditionalSubmitterText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_amendment_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultAmendmentWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_line_numbering"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultLineNumbering)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_sorting"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultSorting)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_default_workflow_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsDefaultWorkflowID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_editor"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableEditor)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_reason_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableReasonOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_recommendation_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableRecommendationOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_sidebox_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableSideboxOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_text_on_projector"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableTextOnProjector)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_enable_working_group_speaker"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsEnableWorkingGroupSpeaker)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_follow_recommendation"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportFollowRecommendation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_submitter_recommendation"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportSubmitterRecommendation)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_export_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsExportTitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_hide_metadata_background"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsHideMetadataBackground)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_line_length"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsLineLength)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_min_digits"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberMinDigits)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_number_with_blank"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsNumberWithBlank)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_preamble"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsPreamble)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_reason_required"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsReasonRequired)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_recommendation_text_mode"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsRecommendationTextMode)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_recommendations_by"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsRecommendationsBy)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_show_referring_motions"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsShowReferringMotions)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_show_sequential_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsShowSequentialNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["motions_supporters_min_amount"]; ok {
		err := json.Unmarshal([]byte(val), &m.MotionsSupportersMinAmount)
		if err != nil {
			return err
		}
	}

	if val, ok := data["name"]; ok {
		err := json.Unmarshal([]byte(val), &m.Name)
		if err != nil {
			return err
		}
	}

	if val, ok := data["option_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OptionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["organization_tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.OrganizationTagIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["personal_note_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PersonalNoteIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["point_of_order_category_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PointOfOrderCategoryIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ballot_paper_number"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollBallotPaperNumber)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ballot_paper_selection"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollBallotPaperSelection)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_candidate_list_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCandidateListIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_countdown_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCountdownID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_couple_countdown"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollCoupleCountdown)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_backend"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultBackend)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_method"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultMethod)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_onehundred_percent_base"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultOnehundredPercentBase)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_default_type"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollDefaultType)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["poll_sort_poll_result_by_votes"]; ok {
		err := json.Unmarshal([]byte(val), &m.PollSortPollResultByVotes)
		if err != nil {
			return err
		}
	}

	if val, ok := data["present_user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.PresentUserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projection_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectionIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_countdown_default_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownDefaultTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_countdown_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_countdown_warning_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorCountdownWarningTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["projector_message_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.ProjectorMessageIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["reference_projector_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.ReferenceProjectorID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["speaker_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.SpeakerIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["start_time"]; ok {
		err := json.Unmarshal([]byte(val), &m.StartTime)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["structure_level_list_of_speakers_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.StructureLevelListOfSpeakersIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["tag_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TagIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["template_for_organization_id"]; ok {
		err := json.Unmarshal([]byte(val), &m.TemplateForOrganizationID)
		if err != nil {
			return err
		}
	}

	if val, ok := data["topic_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TopicIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["topic_poll_default_group_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.TopicPollDefaultGroupIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["user_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.UserIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_allow_self_set_present"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersAllowSelfSetPresent)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_body"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailBody)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_replyto"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailReplyto)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_sender"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSender)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_email_subject"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEmailSubject)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_presence_view"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnablePresenceView)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_vote_delegations"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnableVoteDelegations)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_enable_vote_weight"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersEnableVoteWeight)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_as_submitter"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorAsSubmitter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_as_supporter"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorAsSupporter)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_in_list_of_speakers"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorInListOfSpeakers)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_forbid_delegator_to_vote"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersForbidDelegatorToVote)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_welcometext"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWelcometext)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_welcometitle"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWelcometitle)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_encryption"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanEncryption)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_password"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanPassword)
		if err != nil {
			return err
		}
	}

	if val, ok := data["users_pdf_wlan_ssid"]; ok {
		err := json.Unmarshal([]byte(val), &m.UsersPdfWlanSsid)
		if err != nil {
			return err
		}
	}

	if val, ok := data["vote_ids"]; ok {
		err := json.Unmarshal([]byte(val), &m.VoteIDs)
		if err != nil {
			return err
		}
	}

	if val, ok := data["welcome_text"]; ok {
		err := json.Unmarshal([]byte(val), &m.WelcomeText)
		if err != nil {
			return err
		}
	}

	if val, ok := data["welcome_title"]; ok {
		err := json.Unmarshal([]byte(val), &m.WelcomeTitle)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Meeting) GetRelatedModelsAccessor() *RelatedModelsAccessor {
	return &RelatedModelsAccessor{
		m.GetFqids,
		m.SetRelated,
		m.SetRelatedJSON,
	}
}
