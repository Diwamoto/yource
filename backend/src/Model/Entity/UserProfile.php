<?php
declare(strict_types=1);

namespace App\Model\Entity;

use Cake\ORM\Entity;

/**
 * UserProfile Entity
 *
 * @property int $id
 * @property int $user_id
 * @property string|null $profile
 * @property string|null $birthday
 * @property string|null $from
 * @property string|null $job
 * @property string|null $twitter
 * @property string|null $facebook
 * @property string|null $instagram
 * @property string|null $other
 * @property \Cake\I18n\FrozenTime|null $created
 * @property \Cake\I18n\FrozenTime|null $modified
 *
 * @property \App\Model\Entity\User $user
 */
class UserProfile extends Entity
{
    /**
     * Fields that can be mass assigned using newEntity() or patchEntity().
     *
     * Note that when '*' is set to true, this allows all unspecified fields to
     * be mass assigned. For security purposes, it is advised to set '*' to false
     * (or remove it), and explicitly make individual fields accessible as needed.
     *
     * @var array
     */
    protected $_accessible = [
        'user_id' => true,
        'profile' => true,
        'birthday' => true,
        'from' => true,
        'job' => true,
        'twitter' => true,
        'facebook' => true,
        'instagram' => true,
        'other' => true,
        'created' => true,
        'modified' => true,
        'user' => true,
    ];
}
