package calendar

import (
	"github.com/spf13/cobra"
)

// NewNoteCmd crea el comando padre 'note' y adjunta todos los subcomandos de notas.
func NewCalendarCmd() *cobra.Command {
	var noteCmd = &cobra.Command{
		Use:   "calendar",
		Short: "Gestiona tu calendario",
		Long:  "Permite crear, listar, actualizar y eliminar eventos de Google Calendar.",
	}

	// Adjuntar subcomandos, pasando las dependencias a cada uno
	noteCmd.AddCommand(NewAddCmd())
	noteCmd.AddCommand(NewGetCmd())
	noteCmd.AddCommand(NewListCmd())
	noteCmd.AddCommand(NewUpdateCmd())
	noteCmd.AddCommand(NewDeleteCmd())

	return noteCmd
}

// type Event struct {
// 	AnyoneCanAddSelf bool `json:"anyoneCanAddSelf,omitempty"`
// 	Attachments []*EventAttachment `json:"attachments,omitempty"`
// 	Attendees []*EventAttendee `json:"attendees,omitempty"`
// 	AttendeesOmitted bool `json:"attendeesOmitted,omitempty"`
// 	BirthdayProperties *EventBirthdayProperties `json:"birthdayProperties,omitempty"`
// 	ColorId string `json:"colorId,omitempty"`
// 	ConferenceData *ConferenceData `json:"conferenceData,omitempty"`
// 	Created string `json:"created,omitempty"`
// 	Creator *EventCreator `json:"creator,omitempty"`
// 	Description string `json:"description,omitempty"`
// 	End *EventDateTime `json:"end,omitempty"`
// 	EndTimeUnspecified bool `json:"endTimeUnspecified,omitempty"`
// 	Etag string `json:"etag,omitempty"`
// 	EventType string `json:"eventType,omitempty"`
// 	ExtendedProperties *EventExtendedProperties `json:"extendedProperties,omitempty"`
// 	FocusTimeProperties *EventFocusTimeProperties `json:"focusTimeProperties,omitempty"`
// 	Gadget *EventGadget `json:"gadget,omitempty"`
// 	GuestsCanInviteOthers *bool `json:"guestsCanInviteOthers,omitempty"`
// 	GuestsCanModify bool `json:"guestsCanModify,omitempty"`
// 	GuestsCanSeeOtherGuests *bool `json:"guestsCanSeeOtherGuests,omitempty"`
// 	HangoutLink string `json:"hangoutLink,omitempty"`
// 	HtmlLink string `json:"htmlLink,omitempty"`
// 	ICalUID string `json:"iCalUID,omitempty"`
// 	Id string `json:"id,omitempty"`
// 	Kind string `json:"kind,omitempty"`
// 	Location string `json:"location,omitempty"`
// 	Locked bool `json:"locked,omitempty"`
// 	Organizer *EventOrganizer `json:"organizer,omitempty"`
// 	OriginalStartTime *EventDateTime `json:"originalStartTime,omitempty"`
// 	OutOfOfficeProperties *EventOutOfOfficeProperties `json:"outOfOfficeProperties,omitempty"`
// 	PrivateCopy bool `json:"privateCopy,omitempty"`
// 	Recurrence []string `json:"recurrence,omitempty"`
// 	RecurringEventId string `json:"recurringEventId,omitempty"`
// 	Reminders *EventReminders `json:"reminders,omitempty"`
// 	Sequence int64 `json:"sequence,omitempty"`
// 	Source *EventSource `json:"source,omitempty"`
// 	Start *EventDateTime `json:"start,omitempty"`
// 	Status string `json:"status,omitempty"`
// 	Summary string `json:"summary,omitempty"`
// 	Transparency string `json:"transparency,omitempty"`
// 	Updated string `json:"updated,omitempty"`
// 	Visibility string `json:"visibility,omitempty"`
// 	WorkingLocationProperties *EventWorkingLocationProperties `json:"workingLocationProperties,omitempty"`
// 	googleapi.ServerResponse `json:"-"`
// 	ForceSendFields []string `json:"-"`
// 	NullFields []string `json:"-"`
// }